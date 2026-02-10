package ui

import (
	"Hertz-Hunter-USB-Client/usb"
	"errors"
)

// Refresh list of available serial ports
func (u *Ui) refreshPortsDisplay() {
	availablePorts, err := usb.GetAvailablePorts()
	if err != nil {
		u.showError(err)
		return
	}

	u.portsSelect.Options = availablePorts

	// Show error if no ports found
	if len(availablePorts) == 0 {
		u.showError(errors.New("no serial ports found"))
	}

	u.portsSelect.Refresh()
}
