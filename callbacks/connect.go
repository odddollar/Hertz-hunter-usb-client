package callbacks

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"Hertz-Hunter-USB-Client/usb"
	"errors"

	"fyne.io/fyne/v2"
)

// Attempt to connect to usb serial
func ConnectUSBSerial() {
	// Error if no port selected
	if global.Ui.Ports.SelectedIndex() == -1 {
		dialogs.ShowError(errors.New("port must be selected"))
		return
	}

	// Disable ui elements whilst attempting connection
	global.DisableConnectionUI()

	// Get port
	portName := global.Ui.Ports.Selected

	// Get baud rate
	baudRate := global.Baudrates[global.Ui.Baudrate.SelectedIndex()]

	// Create new connection
	var err error
	global.Connection, err = usb.NewConnection(portName, baudRate)
	if err != nil {
		global.EnableConnectionUI()
		dialogs.ShowError(err)
		return
	}

	dialogs.ShowSuccess("Successfully connected to port")

	// Switch which button is visible
	fyne.Do(func() {
		global.SwitchConnectionButtons()
	})
}
