package callbacks

import (
	"Hertz-Hunter-USB-Client/global"
)

func DisconnectUSBSerial() {
	// Switch which button is visible
	global.SwitchConnectionButtons()

	// Cancel connection
	global.Schema.Stop()

	global.EnableConnectionUI()
}
