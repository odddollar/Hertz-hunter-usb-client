package ui

// Callback for disconnect button
func (u *Ui) disconnectUSBSerial() {
	// Cancel connection
	u.schema.Stop()

	// Switch ui elements
	u.enableConnectionUi()
	u.switchConnectionButtons()
	u.disableSettingsUi()
}
