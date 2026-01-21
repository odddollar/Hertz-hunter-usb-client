package main

import (
	"Hertz-Hunter-USB-Client/global"
	"Hertz-Hunter-USB-Client/usbSerial"
	"fmt"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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

	// Create ports label
	global.Ui.PortsLabel = widget.NewLabel("Serial Port:")

	// Create port selection dropdown with serial ports
	global.Ui.Ports = widget.NewSelect([]string{}, func(selected string) {})

	// Create refresh ports button
	global.Ui.PortsRefresh = widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		usbSerial.RefreshPortsDisplay()
	})

	// Create baudrate label and entry
	global.Ui.BaudrateLabel = widget.NewLabel("Baudrate:")
	global.Ui.Baudrate = widget.NewEntry()
	global.Ui.Baudrate.SetText(fmt.Sprint(global.DefaultBaudrate))

	// Create refresh graph label and dropdown
	global.Ui.GraphRefreshIntervalLabel = widget.NewLabel("Graph Refresh Interval:")
	global.Ui.GraphRefreshInterval = widget.NewSelect([]string{"0.25s", "0.5s", "1s", "2s"}, func(selected string) {})
	global.Ui.GraphRefreshInterval.SetSelected("0.5s")

	// Create connect button
	global.Ui.Connect = widget.NewButton("Connect", func() {})
	global.Ui.Connect.Importance = widget.HighImportance

	// Create graph display area
	global.Ui.Graph = canvas.NewImageFromImage(global.CurrentGraph)

	// Create window layout and set content
	global.W.SetContent(container.NewBorder(
		container.NewVBox(
			container.NewBorder(
				nil,
				nil,
				container.NewVBox(
					global.Ui.PortsLabel,
					global.Ui.BaudrateLabel,
					global.Ui.GraphRefreshIntervalLabel,
				),
				nil,
				container.NewVBox(
					container.NewBorder(
						nil,
						nil,
						nil,
						global.Ui.PortsRefresh,
						global.Ui.Ports,
					),
					global.Ui.Baudrate,
					global.Ui.GraphRefreshInterval,
				),
			),
			global.Ui.Connect,
		),
		nil,
		nil,
		nil,
		global.Ui.Graph,
	))

	// Initial refresh of available ports
	usbSerial.RefreshPortsDisplay()

	// Show window and run app
	global.W.Resize(fyne.NewSize(800, 600))
	global.W.Show()
	global.A.Run()
}
