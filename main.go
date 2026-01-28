package main

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"Hertz-Hunter-USB-Client/usb"
	"Hertz-Hunter-USB-Client/utils"
	"Hertz-Hunter-USB-Client/widgets"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create window
	global.A = app.New()
	global.W = global.A.NewWindow("Hertz Hunter USB Client")

	// Create title widget
	global.Ui.Title = canvas.NewText("Hertz Hunter USB Client", color.Black)
	global.Ui.Title.Alignment = fyne.TextAlignCenter
	global.Ui.Title.TextStyle.Bold = true
	global.Ui.Title.TextSize = 20

	// Create about button
	global.Ui.About = widget.NewButtonWithIcon("", theme.InfoIcon(), dialogs.ShowAbout)

	// Create ports label
	global.Ui.PortsLabel = widget.NewLabel("Serial Port:")

	// Create port selection dropdown with serial ports
	global.Ui.Ports = widget.NewSelect([]string{}, func(s string) {})

	// Create refresh ports button
	global.Ui.PortsRefresh = widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), usb.RefreshPortsDisplay)

	// Create baudrate label and entry
	global.Ui.BaudrateLabel = widget.NewLabel("Baudrate:")
	global.Ui.Baudrate = widget.NewSelect(utils.IntsToStrings(global.Baudrates), func(s string) {})
	global.Ui.Baudrate.SetSelected(fmt.Sprint(global.DefaultBaudrate))

	// Create refresh graph label and dropdown
	global.Ui.GraphRefreshIntervalLabel = widget.NewLabel("Graph Refresh Interval:")
	global.Ui.GraphRefreshInterval = widget.NewSelect(utils.DurationsToStrings(global.RefreshIntervals), func(s string) {})
	global.Ui.GraphRefreshInterval.SetSelected(fmt.Sprintf("%.2gs", global.DefaultRefreshInterval.Seconds()))

	// Create connect button
	global.Ui.Connect = widget.NewButton("Connect", func() {})
	global.Ui.Connect.Importance = widget.HighImportance

	// Create graph display area
	global.Ui.Graph = canvas.NewImageFromImage(global.CurrentGraph)
	global.Ui.Graph.FillMode = canvas.ImageFillStretch  // Pixel perfect scaling
	global.Ui.Graph.ScaleMode = canvas.ImageScalePixels // Nearest neighbor for pixel perfect

	// Create accordion for configuration items
	configAccordion := widget.NewAccordion(widget.NewAccordionItem("Configuration",
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
	))
	configAccordion.Open(0)

	// Create window layout and set content
	global.W.SetContent(container.NewBorder(
		container.NewVBox(
			container.NewBorder(
				nil,
				nil,
				widgets.NewSpacer(widget.NewButtonWithIcon("", theme.InfoIcon(), func() {}).MinSize()), // Keeps title centred
				global.Ui.About,
				global.Ui.Title,
			),
			configAccordion,
		),
		nil,
		nil,
		nil,
		global.Ui.Graph,
	))

	// Initial refresh of available ports
	usb.RefreshPortsDisplay()

	// Show window and run app
	global.W.Resize(fyne.NewSize(800, 600))
	global.W.Show()
	global.A.Run()
}
