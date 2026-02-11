package ui

import (
	"Hertz-Hunter-USB-Client/schema"
	"Hertz-Hunter-USB-Client/utils"
	"errors"

	"fyne.io/fyne/v2"
)

// Attempt to connect to usb serial
func (u *Ui) connectUSBSerial() {
	// Error if no port selected
	if u.portsSelect.SelectedIndex() == -1 {
		u.showError(errors.New("port must be selected"))
		return
	}

	// Disable ui elements whilst attempting connection
	u.disableConnectionUI()

	// Get port
	portName := u.portsSelect.Selected

	// Get baud rate
	baudRate := BAUDRATES[u.baudrateSelect.SelectedIndex()]

	// Get poll rate
	pollRate := REFRESH_INTERVALS[u.graphRefreshIntervalSelect.SelectedIndex()]

	// Create new schema
	var err error
	u.schema, err = schema.NewSchema(portName, baudRate)
	if err != nil {
		u.enableConnectionUI()
		u.showError(err)
		return
	}

	// Switch which button is visible
	u.switchConnectionButtons()

	// Get calibration values
	lowCalibration, highCalibration, err := u.schema.GetCalibratedValues()
	if err != nil {
		// Stop and clear schema
		u.schema.Stop()

		u.enableConnectionUI()
		u.switchConnectionButtons()
		u.showError(err)
		return
	}

	// Start polling for values
	valuesCh, errCh := u.schema.StartPollValues(pollRate)

	go func() {
		for {
			select {
			case values, ok := <-valuesCh: // Update graph
				if !ok {
					return
				}

				img := utils.CreateGraph(
					values.Values,
					GRAPH_WIDTH,
					GRAPH_HEIGHT,
					lowCalibration,
					highCalibration,
				)

				u.currentGraphImage = img
				u.graphImage.Image = u.currentGraphImage

				fyne.Do(func() {
					u.graphImage.Refresh()
					u.switchBandLabels(values.Lowband)
				})

			case err, ok := <-errCh: // Handle errors
				if !ok {
					return
				}

				// Stop and clear schema
				u.schema.Stop()

				// Update ui
				u.enableConnectionUI()
				u.switchConnectionButtons()
				u.showError(err)

				return
			}
		}
	}()
}
