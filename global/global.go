package global

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var (
	Baudrates       = []int{9600, 19200, 38400, 57600, 115200}
	DefaultBaudrate = 115200

	RefreshIntervals       = []time.Duration{250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second}
	DefaultRefreshInterval = 500 * time.Millisecond
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
