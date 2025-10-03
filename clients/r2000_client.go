package clients

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/viniggj2005/r2000-go/connection"
	"github.com/viniggj2005/r2000-go/dtos"
	"github.com/viniggj2005/r2000-go/enums"
	"github.com/viniggj2005/r2000-go/utils"
	"go.bug.st/serial"
)

type R2000Options struct {
	BaudRate    int
	DataBits    int
	Parity      serial.Parity
	StopBits    serial.StopBits
	ReadTimeout time.Duration
	FrameBuf    int
	ErrBuf      int
	Validate    connection.ValidaFrame
}

type R2000Client struct {
	Online   bool
	Name     string
	Port     serial.Port
	FramesCh chan []byte
	ErrorsCh chan error

	Callbacks dtos.OnReadingCallbacks

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	rtMu     sync.Mutex
	rtCancel context.CancelFunc

	writeMu sync.Mutex
}

func (c *R2000Client) GetName() string { return c.Name }

func NewR2000Client(ctx context.Context, name, portName string, cb dtos.OnReadingCallbacks, opts R2000Options) (*R2000Client, error) {
	if opts.FrameBuf <= 0 {
		opts.FrameBuf = 256
	}
	if opts.ErrBuf <= 0 {
		opts.ErrBuf = 16
	}
	ctx, cancel := context.WithCancel(ctx)

	port, err := connection.OpenSerial(ctx, portName, connection.OpenOptions{
		BaudRate:    opts.BaudRate,
		DataBits:    opts.DataBits,
		Parity:      opts.Parity,
		StopBits:    opts.StopBits,
		ReadTimeout: opts.ReadTimeout,
	})
	if err != nil {
		cancel()
		return nil, err
	}

	c := &R2000Client{
		Online:    true,
		Name:      name,
		Port:      port,
		FramesCh:  make(chan []byte, opts.FrameBuf),
		ErrorsCh:  make(chan error, opts.ErrBuf),
		Callbacks: cb,
		ctx:       ctx,
		cancel:    cancel,
	}

	connection.StartReader(ctx, c.Port, c.FramesCh, c.ErrorsCh, opts.Validate)

	c.wg.Add(1)
	go c.processFrames()

	c.wg.Add(1)
	go c.processErrors()

	return c, nil
}

func (c *R2000Client) Close() error {
	c.cancel()
	c.stopRealtimeIfRunning()
	c.wg.Wait()
	return c.Port.Close()
}

func (c *R2000Client) processFrames() {
	defer c.wg.Done()
	for {
		select {
		case <-c.ctx.Done():
			return
		case frame, ok := <-c.FramesCh:
			if !ok {
				return
			}
			if cb := c.Callbacks.OnRawFrame; cb != nil {
				cb(frame)
			}
			// Decoder central já chama os callbacks adequados:
			ProcessR2000Frames(c, frame)
		}
	}
}

func (c *R2000Client) processErrors() {
	defer c.wg.Done()
	for {
		select {
		case <-c.ctx.Done():
			return
		case err, ok := <-c.ErrorsCh:
			if !ok {
				return
			}
			if cb := c.Callbacks.OnError; cb != nil {
				cb(err)
			} else if cb2 := c.Callbacks.OnReadingError; cb2 != nil {
				cb2(c, err.Error())
			}
		}
	}
}

func (c *R2000Client) sendFrame(frame []byte) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	total := 0
	for total < len(frame) {
		n, err := c.Port.Write(frame[total:])
		if err != nil {
			return err
		}
		total += n
	}
	return nil
}

// ===== Comandos =====

func (c *R2000Client) ModuleReset() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.RESET}))
}

func (c *R2000Client) GetModuleTemperature() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_READER_TEMPERATURE}))
}

func (c *R2000Client) GetDrmStatus() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_DRM_STATUS}))
}

func (c *R2000Client) GetFirmwareVersion() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_FIRMWARE_VERSION}))
}

func (c *R2000Client) GetFrequencyRegion() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_FREQUENCY_REGION}))
}

func (c *R2000Client) GetOutPutPower() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_RF_POWER}))
}

func (c *R2000Client) GetWorkAntenna() {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_WORK_ANTENNA}))
}

func (c *R2000Client) SetBeeperMode(mode enums.R2000BeeperEnum) {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{
		Command: enums.SET_BEEPER_MODE,
		Params:  []byte{byte(mode)},
	}))
}

func (c *R2000Client) SetFrequencyRegion(obj dtos.FrequencyRegionsStruct) {
	params := []byte{obj.Region, obj.StartFrequency, obj.EndFrequency}
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.SET_FREQUENCY_REGION, Params: params})
	_ = c.sendFrame(frame)
}

func (c *R2000Client) SetDrm(param enums.R2000StateEnum) {
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{
		Command: enums.SET_DRM,
		Params:  []byte{byte(param)},
	}))
}

func (c *R2000Client) SetOutputPower(dbmPower int) {
	if dbmPower < 20 || dbmPower > 33 {
		if c.Callbacks.OnSetWorkAntenna != nil {
			c.Callbacks.OnSetWorkAntenna(c, false, fmt.Sprintf("valor inválido: %d (use 20–33 dBm)", dbmPower))
		}
		return
	}
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{
		Command: enums.SET_TEMPORARY_OUTPUT_POWER,
		Params:  []byte{byte(dbmPower)},
	}))
}

func (c *R2000Client) SetWorkAntenna(antennaId byte) {
	if antennaId > 0x03 {
		if c.Callbacks.OnError != nil {
			c.Callbacks.OnError(fmt.Errorf("antennaId inválida: %d. Use 0–3", antennaId))
		}
		return
	}
	_ = c.sendFrame(utils.BuildCommandFrame(dtos.BuildFrame{
		Command: enums.SET_WORK_ANTENNA,
		Params:  []byte{antennaId},
	}))
}

// ===== Realtime =====

type RealtimeDto = dtos.RealtimeDto

func (c *R2000Client) StartRealtime(dto *RealtimeDto) error {
	if dto == nil {
		return fmt.Errorf("dto não pode ser nil")
	}

	var antList []byte
	for _, ant := range dto.Antennas {
		if ant > 0x03 {
			return fmt.Errorf("AntennaID inválida: %d", ant)
		}
		antList = append(antList, byte(ant))
	}
	if len(antList) == 0 {
		return fmt.Errorf("forneça pelo menos uma antena")
	}

	c.rtMu.Lock()
	defer c.rtMu.Unlock()
	if c.rtCancel != nil {
		return fmt.Errorf("realtime já em execução")
	}

	rtCtx, rtCancel := context.WithCancel(c.ctx)
	c.rtCancel = rtCancel

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		defer func() {
			c.rtMu.Lock()
			c.rtCancel = nil
			c.rtMu.Unlock()
		}()

		for {
			select {
			case <-rtCtx.Done():
				return
			default:
			}

			for _, ant := range antList {
				select {
				case <-rtCtx.Done():
					return
				default:
				}

				c.GetModuleTemperature()
				c.SetWorkAntenna(ant)

				dwell := time.Duration(dto.DwellS * float64(time.Second))
				timeEnd := time.Now().Add(dwell)
				for time.Now().Before(timeEnd) {
					select {
					case <-rtCtx.Done():
						return
					default:
					}

					inv := utils.BuildCommandFrame(dtos.BuildFrame{
						Command: enums.REAL_TIME_INVENTORY,
						Params:  []byte{byte(dto.Repeat & 0xFF)},
					})
					_ = c.sendFrame(inv)
					time.Sleep(10 * time.Millisecond)
				}

				if dto.SwitchDelayS > 0 {
					select {
					case <-rtCtx.Done():
						return
					case <-time.After(time.Duration(dto.SwitchDelayS * float64(time.Second))):
					}
				}
			}
		}
	}()

	return nil
}

func (c *R2000Client) StopRealtime() {
	c.stopRealtimeIfRunning()
	c.ModuleReset()
}

func (c *R2000Client) stopRealtimeIfRunning() {
	c.rtMu.Lock()
	defer c.rtMu.Unlock()
	if c.rtCancel != nil {
		c.rtCancel()
	}
}
