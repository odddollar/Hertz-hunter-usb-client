package ui

import (
	"Hertz-Hunter-USB-Client/schema"
	"Hertz-Hunter-USB-Client/utils"
	"Hertz-Hunter-USB-Client/widgets"
	"fmt"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	// Main app elements
	a fyne.App
	w fyne.Window

	// Ui components
	titleLabel                 *canvas.Text
	aboutButton                *widget.Button
	portsSelect                *widget.Select
	portsRefreshButton         *widget.Button
	baudrateSelect             *widget.Select
	graphRefreshIntervalSelect *widget.Select
	connectButton              *widget.Button
	disconnectButton           *widget.Button
	graphImage                 *canvas.Image
	leftRssiLabels             *fyne.Container
	rightRssiLabels            *fyne.Container
	highbandFrequencyLabels    *fyne.Container
	lowbandFrequencyLabels     *fyne.Container

	// Store current graph image
	currentGraphImage image.Image

	// Global schema object store
	schema *schema.Schema
}

// Create new ui layout
func (u *Ui) NewUI() {
	// Create window
	u.a = app.New()
	u.w = u.a.NewWindow("Hertz Hunter USB Client")

	// Create title widget
	u.titleLabel = canvas.NewText("Hertz Hunter USB Client", color.Black)
	u.titleLabel.Alignment = fyne.TextAlignCenter
	u.titleLabel.TextStyle.Bold = true
	u.titleLabel.TextSize = 20

	// Create about button
	u.aboutButton = widget.NewButtonWithIcon("", theme.InfoIcon(), u.showAbout)

	// Create port selection dropdown with serial ports
	u.portsSelect = widget.NewSelect([]string{}, func(s string) {})

	// Create refresh ports button
	u.portsRefreshButton = widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), u.refreshPortsDisplay)

	// Create baudrate entry
	u.baudrateSelect = widget.NewSelect(utils.IntsToStrings(BAUDRATES), func(s string) {})
	u.baudrateSelect.SetSelected(fmt.Sprint(DEFAULT_BAUDRATE))

	// Create refresh graph dropdown
	u.graphRefreshIntervalSelect = widget.NewSelect(utils.DurationsToStrings(REFRESH_INTERVALS), func(s string) {})
	u.graphRefreshIntervalSelect.SetSelected(fmt.Sprintf("%.2gs", DEFAULT_REFRESH_INTERVAL.Seconds()))

	// Create connect button
	u.connectButton = widget.NewButton("Connect", func() { go u.connectUSBSerial() })
	u.connectButton.Importance = widget.HighImportance

	// Create disconnect button
	u.disconnectButton = widget.NewButton("Disconnect", u.disconnectUSBSerial)
	u.disconnectButton.Hide()

	// Create graph display area
	u.currentGraphImage = utils.NewEmptyImage(GRAPH_WIDTH, GRAPH_HEIGHT, color.Black)
	u.graphImage = canvas.NewImageFromImage(u.currentGraphImage)
	u.graphImage.FillMode = canvas.ImageFillStretch  // Pixel perfect scaling
	u.graphImage.ScaleMode = canvas.ImageScalePixels // Nearest neighbor for pixel perfect

	// Create rssi labels
	u.leftRssiLabels = newRssiScale(fyne.TextAlignTrailing)
	u.rightRssiLabels = newRssiScale(fyne.TextAlignLeading)

	// Create highband labels
	u.highbandFrequencyLabels = newFrequencyScale("5645MHz", "5795MHz", "5945MHz")

	// Create lowband labels
	u.lowbandFrequencyLabels = newFrequencyScale("5345MHz", "5495MHz", "5645MHz")
	u.lowbandFrequencyLabels.Hide()

	// Create container for connection items
	connectionContainer := container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("Serial Port", container.NewBorder(
				nil,
				nil,
				nil,
				u.portsRefreshButton,
				u.portsSelect,
			)),
			widget.NewFormItem("Baudrate", u.baudrateSelect),
			widget.NewFormItem("Graph Refresh Interval", u.graphRefreshIntervalSelect),
		),
		u.connectButton,
		u.disconnectButton,
	)

	// Create container for calibration items
	calibrationContainer := container.NewVBox()

	// Intermediate accordion so connection can be open by default
	innerAccordion := widget.NewAccordion(
		widget.NewAccordionItem("Connection", connectionContainer),
		widget.NewAccordionItem("Calibration", calibrationContainer),
	)
	innerAccordion.Open(0)

	// Create accordion for configuration items
	configAccordion := widget.NewAccordion(widget.NewAccordionItem("Configuration", innerAccordion))
	configAccordion.Open(0)

	// Create window layout and set content
	u.w.SetContent(container.NewBorder(
		container.NewVBox(
			container.NewBorder(
				nil,
				nil,
				widgets.NewSpacer(widget.NewButtonWithIcon("", theme.InfoIcon(), func() {}).MinSize()), // Keeps title centred
				u.aboutButton,
				u.titleLabel,
			),
			configAccordion,
		),
		container.NewVBox(
			u.highbandFrequencyLabels,
			u.lowbandFrequencyLabels,
		),
		u.leftRssiLabels,
		u.rightRssiLabels,
		u.graphImage,
	))

	// Initial refresh of available ports
	u.refreshPortsDisplay()
}

// Show and run app
func (u *Ui) Run() {
	u.w.Resize(fyne.NewSize(800, 600))
	u.w.Show()
	u.a.Run()
}

// Disable ui elements related to connection
func (u *Ui) disableConnectionUI() {
	fyne.Do(func() {
		u.portsSelect.Disable()
		u.portsRefreshButton.Disable()
		u.baudrateSelect.Disable()
		u.graphRefreshIntervalSelect.Disable()
		u.connectButton.Disable()
	})
}

// Enable ui elements related to connection
func (u *Ui) enableConnectionUI() {
	fyne.Do(func() {
		u.portsSelect.Enable()
		u.portsRefreshButton.Enable()
		u.baudrateSelect.Enable()
		u.graphRefreshIntervalSelect.Enable()
		u.connectButton.Enable()
	})
}

// Switch which connection button is visible
func (u *Ui) switchConnectionButtons() {
	if !u.connectButton.Hidden && u.disconnectButton.Hidden {
		fyne.Do(func() {
			u.connectButton.Hide()
			u.disconnectButton.Show()
		})
	} else if u.connectButton.Hidden && !u.disconnectButton.Hidden {
		fyne.Do(func() {
			u.connectButton.Show()
			u.disconnectButton.Hide()
		})
	}
}

// Switch which band labels are visible
func (u *Ui) switchBandLabels(lowband bool) {
	if lowband {
		fyne.Do(func() {
			u.highbandFrequencyLabels.Hide()
			u.lowbandFrequencyLabels.Show()
		})
	} else {
		fyne.Do(func() {
			u.highbandFrequencyLabels.Show()
			u.lowbandFrequencyLabels.Hide()
		})
	}
}
