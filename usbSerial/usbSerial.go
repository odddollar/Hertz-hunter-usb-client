package usbSerial

import (
	"Hertz-Hunter-USB-Client/global"

	"go.bug.st/serial"
)

// Get list of available serial ports
func GetAvailablePorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {

	}

	return ports
}

// Refresh list of available serial ports
func RefreshPorts() {
	availablePorts := GetAvailablePorts()
	global.Ui.Ports.Options = availablePorts

	// Select first option if only one
	if len(availablePorts) == 1 {
		global.Ui.Ports.SetSelected(availablePorts[0])
	}

	global.Ui.Ports.Refresh()
}
