package widgets

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Generate empty image
func newEmptyImage(width, height int, c color.Color) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// Fill background
	for y := range height {
		for x := range width {
			img.Set(x, y, c)
		}
	}

	return img
}

// Custom widget that displays rssi graph and shows bar data when hovered
type RssiGraph struct {
	widget.BaseWidget

	// Ui elements
	graphCanvas *canvas.Image
	tooltipBg   *canvas.Rectangle
	tooltipText *canvas.Text

	// Constants
	graphWidth  int
	graphHeight int

	// Run time relevant state
	tooltipVisible bool
	rssiValues     []int
	minFrequency   int
	maxFrequency   int
}

// Creates new RssiGraph widget
func NewRssiGraph(graphWidth, graphHeight int) *RssiGraph {
	// Create graph canvas from given image
	graphCanvas := canvas.NewImageFromImage(newEmptyImage(graphWidth, graphHeight, color.Black))
	graphCanvas.FillMode = canvas.ImageFillStretch
	graphCanvas.ScaleMode = canvas.ImageScalePixels

	// Create background
	tooltipBg := canvas.NewRectangle(color.RGBA{R: 32, G: 32, B: 36, A: 235})

	// Create text
	tooltipText := canvas.NewText("", color.White)
	tooltipText.TextSize = 13

	// Hide tooltip by default
	tooltipBg.Hide()
	tooltipText.Hide()

	// Create new object
	graph := &RssiGraph{
		graphCanvas:    graphCanvas,
		tooltipBg:      tooltipBg,
		tooltipText:    tooltipText,
		tooltipVisible: false,
		graphWidth:     graphWidth,
		graphHeight:    graphHeight,
	}

	// Extend base widget and return
	graph.ExtendBaseWidget(graph)
	return graph
}

// Updates tooltip when mouse enters widget
func (r *RssiGraph) MouseIn(event *desktop.MouseEvent) {
	r.updateTooltip(event.Position)
}

// Updates tooltip when mouse moves over widget
func (r *RssiGraph) MouseMoved(event *desktop.MouseEvent) {
	r.updateTooltip(event.Position)
}

// Hides tooltip when mouse leaves widget
func (r *RssiGraph) MouseOut() {
	r.tooltipVisible = false
	r.tooltipBg.Hide()
	r.tooltipText.Hide()
	r.Refresh()
}

// Updates graph image
func (r *RssiGraph) UpdateGraph(numbers []int, minCalibration, maxCalibration int, minFrequency, maxFrequency int) {
	if len(numbers) == 0 {
		return
	}

	// Used for calculating tooltip text
	r.rssiValues = numbers
	r.minFrequency = minFrequency
	r.maxFrequency = maxFrequency

	// Create blank image and calculate values
	img := newEmptyImage(r.graphWidth, r.graphHeight, color.Black)
	barWidth := float64(r.graphWidth) / float64(len(numbers))
	valueRange := float64(maxCalibration - minCalibration)

	for i, value := range numbers {
		// Clamp value to range
		if value < minCalibration {
			value = minCalibration
		}
		if value > maxCalibration {
			value = maxCalibration
		}

		// Normalise value
		normalised := float64(value-minCalibration) / valueRange
		barHeight := int(normalised * float64(r.graphHeight))

		x1 := int(float64(i) * barWidth)
		x2 := int(float64(i+1) * barWidth)
		y1 := r.graphHeight - barHeight

		// Draw bar
		for y := y1; y < r.graphHeight; y++ {
			for x := x1; x < x2 && x < r.graphWidth; x++ {
				img.Set(x, y, theme.Color(theme.ColorNamePrimary))
			}
		}
	}

	r.graphCanvas.Image = img
	r.Refresh()
}

// Updates tooltip position and text
func (r *RssiGraph) updateTooltip(localPos fyne.Position) {
	// Get drawn graph dimensions
	drawSize := r.graphCanvas.Size()
	if drawSize.Width <= 0 || drawSize.Height <= 0 {
		return
	}

	// Check if mouse within graph bounds
	insideX := localPos.X >= 0 && localPos.X < drawSize.Width
	insideY := localPos.Y >= 0 && localPos.Y < drawSize.Height
	if !insideX || !insideY {
		if r.tooltipVisible {
			r.tooltipVisible = false
			r.tooltipBg.Hide()
			r.tooltipText.Hide()
			r.Refresh()
		}
		return
	}

	// Calculate number of bars over from 0
	barCount := len(r.rssiValues)
	displayWidth := int(drawSize.Width)
	mouseX := int(localPos.X)
	barsOver := (mouseX * barCount) / displayWidth

	// Calculate hovered frequency
	step := float64(r.maxFrequency-r.minFrequency) / float64(len(r.rssiValues)-1)
	frequency := int(math.Round(step*float64(barsOver) + float64(r.minFrequency)))

	tooltipText := fmt.Sprintf("%dMHz", frequency)
	if r.tooltipText.Text != tooltipText {
		r.tooltipText.Text = tooltipText
		r.Refresh()
	}

	// Displays mouse coordinates
	// Will be updated later
	// bounds := r.graphCanvas.Image.Bounds()
	// relX := localPos.X / drawSize.Width
	// relY := localPos.Y / drawSize.Height
	// x := int(math.Floor(float64(relX * float32(bounds.Dx()))))
	// y := int(math.Floor(float64(relY * float32(bounds.Dy()))))
	// x = min(max(x, 0), bounds.Dx()-1)
	// y = min(max(y, 0), bounds.Dy()-1)
	// tooltipText := fmt.Sprintf("x: %d, y: %d", x, y)
	// if r.tooltipText.Text != tooltipText {
	// 	r.tooltipText.Text = tooltipText
	// 	r.Refresh()
	// }

	// Get proper tooltip sizing
	padding := float32(6)
	offset := float32(12)
	textSize := r.tooltipText.MinSize()
	bgSize := fyne.NewSize(textSize.Width+padding*2, textSize.Height+padding*2)

	// Put tooltip in bottom right corner of cursor
	tx := localPos.X + offset
	ty := localPos.Y + offset

	// Flip position horizontally
	if tx+bgSize.Width > r.Size().Width {
		tx = localPos.X - bgSize.Width
	}

	// Flip position vertically
	if ty+bgSize.Height > r.Size().Height {
		ty = localPos.Y - bgSize.Height
	}

	// Move tooltip
	r.tooltipBg.Move(fyne.NewPos(tx, ty))
	r.tooltipBg.Resize(bgSize)
	r.tooltipText.Move(fyne.NewPos(tx+padding, ty+padding))

	// Show tooltip
	if !r.tooltipVisible {
		r.tooltipVisible = true
		r.tooltipBg.Show()
		r.tooltipText.Show()
	}
	r.Refresh()
}

// Returns new renderer for RssiGraph
func (r *RssiGraph) CreateRenderer() fyne.WidgetRenderer {
	return &rssiGraphRenderer{rssiGraph: r}
}

// Renderer for RssiGraph widget
type rssiGraphRenderer struct {
	rssiGraph *RssiGraph
}

// Returns minimum size of RssiGraph
func (r *rssiGraphRenderer) MinSize() fyne.Size {
	return fyne.NewSize(160, 120)
}

// Lays out image to fill RssiGraph
func (r *rssiGraphRenderer) Layout(size fyne.Size) {
	r.rssiGraph.graphCanvas.Resize(size)
}

// Refreshes RssiGraph
func (r *rssiGraphRenderer) Refresh() {
	r.rssiGraph.graphCanvas.Refresh()
	r.rssiGraph.tooltipBg.Refresh()
	r.rssiGraph.tooltipText.Refresh()
}

// Returns child widgets of RssiGraph
func (r *rssiGraphRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.rssiGraph.graphCanvas,
		r.rssiGraph.tooltipBg,
		r.rssiGraph.tooltipText,
	}
}

// Does nothing as RssiGraph doesn't hold external resources
func (r *rssiGraphRenderer) Destroy() {}
