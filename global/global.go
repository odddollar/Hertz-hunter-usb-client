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
	Ports       *widget.Select
	PortsRefesh *widget.Button
	Connect     *widget.Button
	Graph       *canvas.Image
}
