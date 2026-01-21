package global

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	DefaultBaudrate = 115200
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
	Baudrate                  *widget.Entry
	GraphRefreshIntervalLabel *widget.Label
	GraphRefreshInterval      *widget.Select
	Connect                   *widget.Button
	Graph                     *canvas.Image
}
