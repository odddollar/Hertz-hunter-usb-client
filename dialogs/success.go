package dialogs

import (
	"Hertz-Hunter-USB-Client/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// Standard dialog to show success
func ShowSuccess(message string) {
	fyne.Do(func() { dialog.ShowInformation("Success", message, global.W) })
}
