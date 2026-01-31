package callbacks

import (
	"Hertz-Hunter-USB-Client/global"
)

func DisconnectUSBSerial() {
	// Switch which button is visible
	global.SwitchConnectionButtons()
	global.EnableConnectionUI()

	// Cancel connection
	if global.Schema != nil {
		global.Schema.Stop()
		global.Schema = nil
	}
}
