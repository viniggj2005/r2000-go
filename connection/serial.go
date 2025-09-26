package connection

import (
	"fmt"

	"github.com/viniggj2005/r2000-go/logs"
	"go.bug.st/serial"
)

type PortError struct {
	Code int
}

func GetPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		logs.Error(err.Error())
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
		for {
			n, err := portHandle.Read(buf)
			if err != nil {
				return
			}
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				callback(data)
			}
		}
	}()
}
