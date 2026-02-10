package ui

// Callback for disconnect button
func (u *Ui) disconnectUSBSerial() {
	// Switch which button is visible
	u.switchConnectionButtons()
	u.enableConnectionUI()

	// Cancel connection
	u.schema.Stop()
}
