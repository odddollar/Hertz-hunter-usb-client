package callbacks

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"Hertz-Hunter-USB-Client/schema"
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

	// Get poll rate
	pollRate := global.RefreshIntervals[global.Ui.GraphRefreshInterval.SelectedIndex()]

	// Create new schema
	var err error
	global.Schema, err = schema.NewSchema(portName, baudRate)
	if err != nil {
		global.EnableConnectionUI()
		dialogs.ShowError(err)
		return
	}

	// Switch which button is visible
	fyne.Do(func() {
		global.SwitchConnectionButtons()
	})

	// Start polling for values
	errCh := global.Schema.StartPollValues(pollRate)
	if err = <-errCh; err != nil {
		global.EnableConnectionUI()
		dialogs.ShowError(err)
	}
}
