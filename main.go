package main

import "Hertz-Hunter-USB-Client/ui"

func main() {
	// Create and run ui
	ui := ui.Ui{}
	ui.NewUI()
	ui.Run()
}
