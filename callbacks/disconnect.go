package callbacks

import (
	"Hertz-Hunter-USB-Client/global"
)

func DisconnectUSBSerial() {
	// Switch which button is visible
	global.SwitchConnectionButtons()

	// Cancel polling
	global.Schema.StopPollValues()

	global.Connection.Disconnect()
	global.EnableConnectionUI()
}
