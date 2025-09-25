package serial

import (
	"fmt"

	"go.bug.st/serial"
)

func PortsList() {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ports)
}
