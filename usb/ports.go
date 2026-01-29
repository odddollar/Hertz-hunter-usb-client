package usb

import (
	"go.bug.st/serial"
)

// Get list of available serial ports
func GetAvailablePorts() ([]string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}

	return ports, nil
}
