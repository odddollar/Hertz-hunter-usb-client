package graph

import (
	"image"
	"image/color"
)

// Generates black and white histogram image with one bar per number
func CreateGraph(numbers []int, width, height int) image.Image {
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Fill background with white
	for y := range height {
		for x := range width {
			img.Set(x, y, color.White)
		}
	}

	if len(numbers) == 0 {
		return img
	}

	// Draw histogram bars
	barWidth := float64(width) / float64(len(numbers))
	maxValue := 100.0 // Numbers are 0-100

	for i, value := range numbers {
		// Calculate bar dimensions
		barHeight := int(float64(value) / maxValue * float64(height))
		x1 := int(float64(i) * barWidth)
		x2 := int(float64(i+1) * barWidth)
		y1 := height - barHeight
		y2 := height

		// Draw bar
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				if x < width {
					img.Set(x, y, color.Black)
				}
			}
		}
	}

	return img
}
