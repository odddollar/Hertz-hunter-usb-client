package graph

import (
	"Hertz-Hunter-USB-Client/global"
	"image"
	"image/color"
)

// Generates black and white histogram image with one bar per number
func CreateGraph(numbers []int) image.Image {
	img := image.NewGray(image.Rect(0, 0, global.GraphWidth, global.GraphHeight))

	// Fill background with white
	for y := range global.GraphHeight {
		for x := range global.GraphWidth {
			img.Set(x, y, color.White)
		}
	}

	if len(numbers) == 0 {
		return img
	}

	// Draw histogram bars
	barWidth := float64(global.GraphWidth) / float64(len(numbers))
	maxValue := 100.0 // Numbers are 0-100

	for i, value := range numbers {
		// Calculate bar dimensions
		barHeight := int(float64(value) / maxValue * float64(global.GraphHeight))
		x1 := int(float64(i) * barWidth)
		x2 := int(float64(i+1) * barWidth)
		y1 := global.GraphHeight - barHeight
		y2 := global.GraphHeight

		// Draw bar
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				if x < global.GraphWidth {
					img.Set(x, y, color.Black)
				}
			}
		}
	}

	return img
}
