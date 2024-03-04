package serial

import (
	"go.bug.st/serial"
)

func GetSerialPorts() ([]string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}

	return ports, nil
}
