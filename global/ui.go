package global

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Track current graph image
var CurrentGraph image.Image

// Main app elements
var (
	A fyne.App
	W fyne.Window
)

// Ui components
var Ui struct {
	Title                     *canvas.Text
	About                     *widget.Button
	PortsLabel                *widget.Label
	Ports                     *widget.Select
	PortsRefresh              *widget.Button
	BaudrateLabel             *widget.Label
	Baudrate                  *widget.Select
	GraphRefreshIntervalLabel *widget.Label
	GraphRefreshInterval      *widget.Select
	Connect                   *widget.Button
	Graph                     *canvas.Image
}

// Disable ui elements related to connection
func DisableConnectionUI() {
	fyne.Do(func() {
		Ui.Ports.Disable()
		Ui.PortsRefresh.Disable()
		Ui.Baudrate.Disable()
		Ui.Connect.Disable()
	})
}

// Enable ui elements related to connection
func EnableConnectionUI() {
	fyne.Do(func() {
		Ui.Ports.Enable()
		Ui.PortsRefresh.Enable()
		Ui.Baudrate.Enable()
		Ui.Connect.Enable()
	})
}
