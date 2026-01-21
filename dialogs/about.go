package dialogs

import (
	"Hertz-Hunter-USB-Client/global"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/x/fyne/dialog"
)

// Use Fyne-X extensions to create about window
func ShowAbout() {
	// Parse urls for documentation
	HHDocumentation, _ := url.Parse("https://github.com/odddollar/Hertz-hunter")
	HHClientDocumentation, _ := url.Parse("https://example.com")

	links := []*widget.Hyperlink{
		widget.NewHyperlink("Hertz Hunter Documentation", HHDocumentation),
		widget.NewHyperlink("Hertz Hunter USB Client Documentation", HHClientDocumentation),
	}

	// Markdown program description
	content := "USB serial client for the **Hertz Hunter** spectrum analyser"

	// Use Fyne-X's about dialog
	d := dialog.NewAbout(content, links, global.A, global.W)
	d.Resize(fyne.NewSize(0, 402)) // 0 width as will always be at least content's min-width
	d.Show()
}
