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
	u.disableConnectionUi()

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
		u.enableConnectionUi()
		u.showError(err)
		return
	}

	// Switch which connection button is visible
	u.switchConnectionButtons()

	// Enable settings ui
	u.enableSettingsUi()

	// Get calibration values
	lowCalibration, highCalibration, err := u.schema.GetCalibratedValues()
	if err != nil {
		u.connectionError(err)
		return
	}

	// Update entries with calibration values
	u.updateCalibrationEntries(highCalibration, lowCalibration)

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

				u.connectionError(err)
				return
			}
		}
	}()
}
