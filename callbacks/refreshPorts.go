package callbacks

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"Hertz-Hunter-USB-Client/usb"
	"errors"
)

// Refresh list of available serial ports
func RefreshPortsDisplay() {
	availablePorts, err := usb.GetAvailablePorts()
	if err != nil {
		dialogs.ShowError(err)
		return
	}

	global.Ui.Ports.Options = availablePorts

	// Show error if no ports found
	if len(availablePorts) == 0 {
		dialogs.ShowError(errors.New("no serial ports found"))
	}

	global.Ui.Ports.Refresh()
}
