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
	Disconnect                *widget.Button
	Graph                     *canvas.Image
	HighbandFrequencyLabels   *fyne.Container
	LowbandFrequencyLabels    *fyne.Container
}

// Disable ui elements related to connection
func DisableConnectionUI() {
	fyne.Do(func() {
		Ui.Ports.Disable()
		Ui.PortsRefresh.Disable()
		Ui.Baudrate.Disable()
		Ui.GraphRefreshInterval.Disable()
		Ui.Connect.Disable()
	})
}

// Enable ui elements related to connection
func EnableConnectionUI() {
	fyne.Do(func() {
		Ui.Ports.Enable()
		Ui.PortsRefresh.Enable()
		Ui.Baudrate.Enable()
		Ui.GraphRefreshInterval.Enable()
		Ui.Connect.Enable()
	})
}

// Switch which connection button is visible
func SwitchConnectionButtons() {
	if !Ui.Connect.Hidden && Ui.Disconnect.Hidden {
		fyne.Do(func() {
			Ui.Connect.Hide()
			Ui.Disconnect.Show()
		})
	} else if Ui.Connect.Hidden && !Ui.Disconnect.Hidden {
		fyne.Do(func() {
			Ui.Connect.Show()
			Ui.Disconnect.Hide()
		})
	}
}

// Switch which band labels are visible
func SwitchBandLabels(highband bool) {
	if highband {
		fyne.Do(func() {
			Ui.HighbandFrequencyLabels.Show()
			Ui.LowbandFrequencyLabels.Hide()
		})
	} else {
		fyne.Do(func() {
			Ui.HighbandFrequencyLabels.Hide()
			Ui.LowbandFrequencyLabels.Show()
		})
	}
}
