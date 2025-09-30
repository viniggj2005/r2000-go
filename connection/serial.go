package connection

import (
	"fmt"

	"go.bug.st/serial"
)

type PortError struct {
	Code int
}

func GetPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println(err.Error())
	}
	return ports
}

func OpenSerialConnection(port string) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
		DataBits: 8,
	}
	portHandle, err := serial.Open(port, mode)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir porta %s: %w", port, err)
	}
	return portHandle, nil
}

func ListenSerial(portHandle serial.Port, callback func([]byte)) {
	go func() {
		buf := make([]byte, 256)
		var buffer []byte

		for {
			n, err := portHandle.Read(buf)
			if err != nil {
				return
			}
			if n > 0 {
				buffer = append(buffer, buf[:n]...)

				for {
					if len(buffer) < 2 {
						break
					}
					if buffer[0] != 0xA0 {
						buffer = buffer[1:]
						continue
					}

					length := int(buffer[1])
					frameSize := length + 2
					if len(buffer) < frameSize {
						break
					}

					frame := buffer[:frameSize]
					buffer = buffer[frameSize:]

					callback(frame)
				}
			}
		}
	}()
}
