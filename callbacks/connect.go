package callbacks

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"Hertz-Hunter-USB-Client/graph"
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
	fyne.Do(func() { global.SwitchConnectionButtons() })

	// Get calibration values
	lowCalibration, highCalibration, err := global.Schema.GetCalibratedValues()
	if err != nil {
		// Stop and clear schema
		global.Schema.Stop()
		global.Schema = nil

		global.EnableConnectionUI()
		global.SwitchConnectionButtons()
		dialogs.ShowError(err)
		return
	}

	// Start polling for values
	valuesCh, errCh := global.Schema.StartPollValues(pollRate)

	go func() {
		for {
			select {
			case values, ok := <-valuesCh: // Update graph
				if !ok {
					return
				}

				img := graph.CreateGraph(
					values.Values,
					global.GraphWidth,
					global.GraphHeight,
					lowCalibration,
					highCalibration,
				)

				global.CurrentGraph = img
				global.Ui.Graph.Image = global.CurrentGraph

				fyne.Do(func() {
					global.Ui.Graph.Refresh()
					global.SwitchBandLabels(values.Lowband)
				})

			case err, ok := <-errCh: // Handle errors
				if !ok {
					return
				}

				// Stop and clear schema
				global.Schema.Stop()
				global.Schema = nil

				// Update ui
				global.EnableConnectionUI()
				global.SwitchConnectionButtons()
				dialogs.ShowError(err)

				return
			}
		}
	}()
}
