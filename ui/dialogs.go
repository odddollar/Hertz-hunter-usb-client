package ui

import (
	"Hertz-Hunter-USB-client/widgets"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Use Fyne-X extensions to create about window
func (u *Ui) showAbout() {
	// Parse urls for documentation
	documentation, _ := url.Parse("https://github.com/odddollar/Hertz-Hunter")
	clientDocumentation, _ := url.Parse("https://github.com/odddollar/Hertz-Hunter-USB-client")

	links := []*widget.Hyperlink{
		widget.NewHyperlink("Hertz Hunter Documentation", documentation),
		widget.NewHyperlink("Hertz Hunter USB Client Documentation", clientDocumentation),
	}

	// Markdown program description
	content := "USB serial client for the **Hertz Hunter** spectrum analyser"

	// Use Fyne-X's about dialog
	d := widgets.NewAbout(content, links, u.batteryVoltage, u.a, u.w)
	d.Resize(fyne.NewSize(0, 425)) // 0 width as will always be at least content's min-width
	d.Show()
}

// Show help dialog for communication retries
func (u *Ui) showMaxComRetriesHelp() {
	d := dialog.NewInformation(
		"Max Communication Retries Help",
		"How many times a serial message should attempt to be sent before returning an error",
		u.w,
	)
	d.Resize(fyne.NewSize(400, 0))
	d.Show()
}

// Show help dialog for graph refresh interval
func (u *Ui) showGraphRefreshIntervalHelp() {
	d := dialog.NewInformation(
		"Graph Refresh Interval Help",
		"The delay between polls for updated RSSI data",
		u.w,
	)
	d.Show()
}

// Standard dialog to show error
func (u *Ui) showError(err error) {
	fyne.Do(func() { dialog.ShowError(err, u.w) })
}

// Standard dialog to show success
func (u *Ui) showSuccess(message string) {
	fyne.Do(func() { dialog.ShowInformation("Success", message, u.w) })
}
