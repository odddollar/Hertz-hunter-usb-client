package usbSerial

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"errors"

	"go.bug.st/serial"
)

// Get list of available serial ports
func getAvailablePorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		dialogs.ShowError(err)
	}

	return ports
}

// Refresh list of available serial ports
func RefreshPortsDisplay() {
	availablePorts := getAvailablePorts()
	global.Ui.Ports.Options = availablePorts

	// Show error if no ports found
	if len(availablePorts) == 0 {
		dialogs.ShowError(errors.New("no serial ports found"))
	}

	global.Ui.Ports.Refresh()
}
