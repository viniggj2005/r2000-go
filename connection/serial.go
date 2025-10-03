package connection

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"go.bug.st/serial"
)

const (
	sofByte     = 0xA0
	maxFrameLen = 1024
	defBaud     = 115200
)

type ValidaFrame func([]byte) bool

type OpenOptions struct {
	BaudRate    int
	DataBits    int
	Parity      serial.Parity
	StopBits    serial.StopBits
	ReadTimeout time.Duration
}

func GetPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var out []string
	for _, p := range ports {
		if strings.HasPrefix(p, "/dev/ttyUSB") || strings.HasPrefix(p, "COM") {
			out = append(out, p)
		}
	}
	return out
}

func OpenSerial(ctx context.Context, port string, opt OpenOptions) (serial.Port, error) {
	if opt.BaudRate == 0 {
		opt.BaudRate = defBaud
	}
	if opt.DataBits == 0 {
		opt.DataBits = 8
	}
	if opt.Parity == 0 {
		opt.Parity = serial.NoParity
	}
	if opt.StopBits == 0 {
		opt.StopBits = serial.OneStopBit
	}

	mode := &serial.Mode{
		BaudRate: opt.BaudRate,
		Parity:   opt.Parity,
		StopBits: opt.StopBits,
		DataBits: opt.DataBits,
	}

	type res struct {
		p   serial.Port
		err error
	}
	ch := make(chan res, 1)
	go func() {
		p, err := serial.Open(port, mode)
		ch <- res{p: p, err: err}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r := <-ch:
		if r.err != nil {
			return nil, fmt.Errorf("erro ao abrir porta %s: %w", port, r.err)
		}
		if opt.ReadTimeout <= 0 {
			opt.ReadTimeout = 500 * time.Millisecond
		}
		_ = r.p.SetReadTimeout(opt.ReadTimeout)
		return r.p, nil
	}
}

// StartReader publica frames válidos em framesCh e erros em errCh.
// Frame: [0xA0][LEN][LEN bytes]
func StartReader(ctx context.Context, port serial.Port, framesCh chan<- []byte, errCh chan<- error, validar ValidaFrame) {
	go func() {
		defer func() {
			close(framesCh)
			close(errCh)
		}()

		tmp := make([]byte, 512)
		buf := make([]byte, 0, 4096)

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			n, err := port.Read(tmp)
			if n > 0 {
				buf = append(buf, tmp[:n]...)

				for {
					idx := bytes.IndexByte(buf, sofByte)
					if idx == -1 {
						buf = buf[:0]
						break
					}
					if idx > 0 {
						copy(buf, buf[idx:])
						buf = buf[:len(buf)-idx]
					}
					if len(buf) < 2 {
						break
					}
					L := int(buf[1])
					if L <= 0 || L > maxFrameLen {
						buf = buf[1:]
						continue
					}
					total := 2 + L
					if len(buf) < total {
						break
					}
					frame := make([]byte, total)
					copy(frame, buf[:total])
					copy(buf, buf[total:])
					buf = buf[:len(buf)-total]

					if validar != nil && !validar(frame) {
						continue
					}

					select {
					case framesCh <- frame:
					case <-ctx.Done():
						return
					}
				}
			}

			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				select {
				case errCh <- fmt.Errorf("serial read: %w", err):
				default:
				}
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
}
