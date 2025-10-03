package clients

import (
	"fmt"
	"sync"
	"time"

	"github.com/viniggj2005/r2000-go/connection"
	"github.com/viniggj2005/r2000-go/dtos"
	"github.com/viniggj2005/r2000-go/enums"
	"github.com/viniggj2005/r2000-go/utils"
	"go.bug.st/serial"
)

// Estrutura principal do cliente.
type R2000Client struct {
	Online       bool           // Status de conexão
	Name         string         // Nome do cliente (identificação lógica)
	Port         serial.Port    // Porta serial aberta
	FrameQueue   chan []byte    // Canal para enfileirar frames recebidos
	stopChan     chan struct{}  // Canal usado para sinalizar parada
	wg           sync.WaitGroup // WaitGroup para aguardar goroutines terminarem
	realtimeStop chan struct{}
	realtimeRun  bool
	Callbacks    dtos.OnReadingCallbacks // Callback para processar cada frame recebido
	Watcher      *TimeWatcher
	writeMu      sync.Mutex
}
type TimeWatcher struct {
	TimeWatcherRun      bool
	TimeWatcherStopChan chan struct{}
}

func (c *R2000Client) GetName() string {
	return c.Name
}

// Construtor do cliente: abre a porta, inicia listener e processador.
func NewR2000Client(name, portName string, onFrame dtos.OnReadingCallbacks) (*R2000Client, error) {
	// Abre a porta serial com a função já existente no seu pacote connection
	portHandle, err := connection.OpenSerialConnection(portName)
	if err != nil {
		return nil, err
	}

	// Instancia o cliente
	client := &R2000Client{
		Name:       name,
		Port:       portHandle,
		Online:     true,
		FrameQueue: make(chan []byte, 100), // canal bufferizado para frames
		Callbacks:  onFrame,                // callback fornecido externamente
		stopChan:   make(chan struct{}),    // inicializa canal de parada
	}

	// Usa a função ListenSerial já criada no seu pacote connection
	// Cada vez que chegam bytes, eles são jogados na FrameQueue
	connection.ListenSerial(portHandle, func(data []byte) {
		select {
		case client.FrameQueue <- data: // envia para fila
		default: // se fila cheia, descarta
		}
	})

	// Inicia a goroutine que processa os frames
	client.wg.Add(1) // diz ao waiting group que tem uma goroutine rodando.
	go client.processFrames()

	return client, nil
}

// Fecha o cliente: sinaliza stopChan, aguarda goroutines e fecha a porta.
func (c *R2000Client) Close() error {
	close(c.stopChan)     // avisa para as goroutines pararem
	c.wg.Wait()           // espera todas terminarem
	return c.Port.Close() // fecha a porta serial
}

// Goroutine que consome frames da fila e chama o callback OnFrame.
func (c *R2000Client) processFrames() {
	defer c.wg.Done() //caso o processamento de frame para devido a stopChan ele informa o waitingGroup que uma goroutine terminou.
	for {
		select {
		case <-c.stopChan: //aqui se o stopChan for fechado da maneira close(stopchan) e le para de executar o loop da função de processamento de frame.
			return
		case frame := <-c.FrameQueue: //aqui estou esvaziando a fila e mandando para a variavel frame.
			// >>> aqui estou mandado o que a variavel frame recebeu para o dispatcher para ele processar a resposta
			ProcessR2000Frames(c, frame)
		}
	}
}

// Função responsável por eviar os frames através do canal serial.
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

// Reseta o módulo.
func (c *R2000Client) ModuleReset() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.RESET})
	c.sendFrame(frame)
}

// Retorna a temperatura do módulo.
func (c *R2000Client) GetModuleTemperature() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_READER_TEMPERATURE})
	c.sendFrame(frame)
}

// Retorna o status do Drm.
func (c *R2000Client) GetDrmStatus() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_DRM_STATUS})
	c.sendFrame(frame)
}

// Retorna a versão do firmware.
func (c *R2000Client) GetFirmwareVersion() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_FIRMWARE_VERSION})
	c.sendFrame(frame)
}

// Retorna a Região e a frequencia de operação.
func (c *R2000Client) GetFrequencyRegion() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_FREQUENCY_REGION})
	c.sendFrame(frame)
}

// Retorna a potencuia que o modulo ta entregando.
func (c *R2000Client) GetOutPutPower() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_RF_POWER})
	c.sendFrame(frame)
}

// Retorna a antena em funcionamento no momento.
func (c *R2000Client) GetWorkAntenna() {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.GET_WORK_ANTENNA})
	c.sendFrame(frame)
}

// Altera o comportamento do buzzer.
func (c *R2000Client) SetBeeperMode(mode enums.R2000BeeperEnum) {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.SET_BEEPER_MODE, Params: []byte{byte(mode)}})
	c.sendFrame(frame)
}

// Altera a Região de funcionamento do módulo.
func (c *R2000Client) SetFrequencyRegion(obj dtos.FrequencyRegionsStruct) {
	params := []byte{obj.Region, obj.StartFrequency, obj.EndFrequency}
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.SET_FREQUENCY_REGION, Params: params})
	c.sendFrame(frame)
}

// Altera o estado do Drm.
func (c *R2000Client) SetDrm(param enums.R2000StateEnum) {
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.SET_DRM, Params: []byte{byte(param)}})
	c.sendFrame(frame)
}

// Altera a potência de saída do módulo.
func (c *R2000Client) SetOutputPower(dbmPower int) {
	if dbmPower < 20 || dbmPower > 33 {
		c.Callbacks.OnSetWorkAntenna(c, false, fmt.Sprintf("valor inválido: %d (o valor deve estar 20–33 dBm)", dbmPower))
		return
	}
	frame := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.SET_TEMPORARY_OUTPUT_POWER, Params: []byte{byte(dbmPower)}})
	c.sendFrame(frame)
}

// Altera a antena de saída do módulo.
func (c *R2000Client) SetWorkAntenna(antennaId byte) {
	if antennaId > 0x03 {
		fmt.Printf("valor inválido: %d. Range válido é 0–3\n", antennaId)
		return
	}
	frame := utils.BuildCommandFrame(dtos.BuildFrame{
		Command: enums.SET_WORK_ANTENNA,
		Params:  []byte{antennaId},
	})
	c.sendFrame(frame)
}

// Inicia inventário em loop até chamar StopRealtime().
func (c *R2000Client) StartRealtime(dto *dtos.RealtimeDto) error {
	if c.realtimeRun {
		return fmt.Errorf("realtime já em execução")
	}
	if dto == nil {
		return fmt.Errorf("dto não pode ser nil")
	}

	// valida antenas
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

	c.realtimeStop = make(chan struct{})
	c.realtimeRun = true

	go func() {
		defer func() { c.realtimeRun = false }()

		for {
			select {
			case <-c.realtimeStop:
				return
			default:
			}

			for _, ant := range antList {
				// usa a função de trocar antena
				c.SetWorkAntenna(ant)

				// dwell (tempo que fica naquela antena)
				timeEnd := time.Now().Add(time.Duration(dto.DwellS * float64(time.Second)))
				for time.Now().Before(timeEnd) {
					select {
					case <-c.realtimeStop:
						return
					default:
					}

					inv := utils.BuildCommandFrame(dtos.BuildFrame{
						Command: enums.REAL_TIME_INVENTORY,
						Params:  []byte{byte(dto.Repeat & 0xFF)},
					})
					c.sendFrame(inv)

					// dá tempo para o hardware processar
					time.Sleep(10 * time.Millisecond)
				}

				// delay entre antenas
				if dto.SwitchDelayS > 0 {
					time.Sleep(time.Duration(dto.SwitchDelayS * float64(time.Second)))
				}
			}
		}
	}()

	return nil
}

// Para o inventário realtime.
func (c *R2000Client) StopRealtime() {
	if !c.realtimeRun {
		return
	}
	close(c.realtimeStop)
	c.ModuleReset()
}

func (tw *TimeWatcher) StartTemperatureWatcher(clients []*R2000Client, intervalSeconds int) {
	if tw.TimeWatcherStopChan == nil {
		tw.TimeWatcherStopChan = make(chan struct{})
	}
	tw.TimeWatcherRun = true

	go func() {
		defer func() { tw.TimeWatcherRun = false }()
		ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
		defer ticker.Stop()

		for _, c := range clients {
			c.GetModuleTemperature()
		}

		for {
			select {
			case <-tw.TimeWatcherStopChan:
				return
			case <-ticker.C:
				for _, c := range clients {
					c.GetModuleTemperature()
				}
			}
		}
	}()
}

func (tw *TimeWatcher) StopTemperatureWatcher() {
	if tw.TimeWatcherStopChan != nil {
		close(tw.TimeWatcherStopChan)
		tw.TimeWatcherStopChan = nil
	}
}
