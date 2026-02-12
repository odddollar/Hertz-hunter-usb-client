package ui

// Perform actions required on connection error
func (u *Ui) connectionError(err error) {
	// Stop and clear schema
	u.schema.Stop()

	u.enableConnectionUi()
	u.switchConnectionButtons()
	u.disableSettingsUi()
	u.showError(err)
}
