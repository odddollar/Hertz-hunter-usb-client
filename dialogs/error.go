package dialogs

import (
	"Hertz-Hunter-USB-Client/global"

	"fyne.io/fyne/v2/dialog"
)

// Standard dialog to show error
func ShowError(err error) {
	dialog.ShowError(err, global.W)
}
