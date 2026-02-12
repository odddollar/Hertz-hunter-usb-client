package ui

import (
	"strconv"
)

// Callback for set calibration button
func (u *Ui) setCalibrationValues() {
	// Get values from entries
	high, _ := strconv.Atoi(u.highRssiCalibrationEntry.Text)
	low, _ := strconv.Atoi(u.lowRssiCalibrationEntry.Text)

	// Set values
	err := u.schema.SetCalibratedValues(low, high)
	if err != nil {
		u.showError(err)
	}
}
