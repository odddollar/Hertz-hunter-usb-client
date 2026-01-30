package callbacks

import (
	"Hertz-Hunter-USB-Client/global"
)

func DisconnectUSBSerial() {
	// Switch which button is visible
	global.SwitchConnectionButtons()

	// Cancel polling
	global.Connection.StopPollValues()

	global.Connection.Disconnect()
	global.EnableConnectionUI()
}
