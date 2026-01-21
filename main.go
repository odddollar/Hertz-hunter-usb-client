package main

import (
	"Hertz-Hunter-USB-Client/global"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a blue image for testing
	width, height := 400, 300
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := color.RGBA{0, 0, 255, 255} // Blue color

	// Fill the entire image with blue
	for y := range height {
		for x := range width {
			img.Set(x, y, blue)
		}
	}
	global.CurrentGraph = img

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

	// Create graph display area
	global.Ui.Graph = canvas.NewImageFromImage(global.CurrentGraph)

	// Create window layout and set content
	global.W.SetContent(container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			nil,
			global.Ui.Connect,
			global.Ui.Ports,
		),
		nil,
		nil,
		nil,
		global.Ui.Graph,
	))

	// Show window and run app
	global.W.Resize(fyne.NewSize(800, 600))
	global.W.Show()
	global.A.Run()
}
