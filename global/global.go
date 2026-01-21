package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Main app elements
var (
	A fyne.App
	W fyne.Window
)

// Ui components
var Ui struct {
	Ports   *widget.Select
	Connect *widget.Button
}
