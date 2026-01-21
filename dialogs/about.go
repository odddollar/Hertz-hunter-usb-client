package dialogs

import (
	"Hertz-Hunter-USB-Client/global"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowAbout() {
	// Create layout
	// Separate markdown widget for better spacing
	d := container.NewVBox(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("Version: **%s**", global.A.Metadata().Version)),
		widget.NewRichTextFromMarkdown("Client for USB communication with a Hertz Hunter device"),
		widget.NewRichTextFromMarkdown("Source: [example.com](https://example.com)"),
	)

	// Show information dialog with layout
	dialog.ShowCustom(
		"About",
		"OK",
		d,
		global.W,
	)
}
