package main

import (
	"Hertz-Hunter-USB-client/ui"
)

func main() {
	// Create and run ui
	ui := ui.Ui{}
	ui.NewUI()
	ui.Run()
}
