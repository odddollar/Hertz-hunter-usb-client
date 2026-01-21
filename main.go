package main

import (
	"Hertz-Hunter-USB-Client/global"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create window
	global.A = app.New()
	global.W = global.A.NewWindow("Hertz Hunter USB Client")

	// Create port selection dropdown
	global.Ui.Ports = widget.NewSelect([]string{"COM1", "COM2", "COM3"}, func(selected string) {
		// Handle port selection
	})

	// Create connect button
	global.Ui.Connect = widget.NewButton("Connect", func() {

	})

	// Create window layout and set content
	global.W.SetContent(container.NewBorder(
		nil,
		nil,
		nil,
		global.Ui.Connect,
		global.Ui.Ports,
	))

	// Show window and run app
	global.W.Show()
	global.A.Run()
}
