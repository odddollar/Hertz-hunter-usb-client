package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

// Create rssi scale with given text alignment
func newRssiScale(alignment fyne.TextAlign) *fyne.Container {
	full := canvas.NewText("100%", theme.Color(theme.ColorNameForeground))
	full.Alignment = alignment
	full.TextStyle.Bold = true

	mid := canvas.NewText("50%", theme.Color(theme.ColorNameForeground))
	mid.Alignment = alignment
	mid.TextStyle.Bold = true

	none := canvas.NewText("0%", theme.Color(theme.ColorNameForeground))
	none.Alignment = alignment
	none.TextStyle.Bold = true

	return container.NewBorder(
		full,
		none,
		nil,
		nil,
		mid,
	)
}
