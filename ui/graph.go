package ui

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2/theme"
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

// Generates histogram image with one bar per number
func createGraph(numbers []int, width, height int, minValue, maxValue int) image.Image {
	img := newEmptyImage(width, height, color.Black)

	if len(numbers) == 0 || maxValue <= minValue {
		return img
	}

	barWidth := float64(width) / float64(len(numbers))
	valueRange := float64(maxValue - minValue)

	for i, value := range numbers {
		// Clamp value to range
		if value < minValue {
			value = minValue
		}
		if value > maxValue {
			value = maxValue
		}

		// Normalise value
		normalised := float64(value-minValue) / valueRange
		barHeight := int(normalised * float64(height))

		x1 := int(float64(i) * barWidth)
		x2 := int(float64(i+1) * barWidth)
		y1 := height - barHeight

		// Draw bar
		for y := y1; y < height; y++ {
			for x := x1; x < x2 && x < width; x++ {
				img.Set(x, y, theme.Color(theme.ColorNamePrimary))
			}
		}
	}

	return img
}
