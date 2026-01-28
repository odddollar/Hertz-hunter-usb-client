package usb

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"bufio"
	"errors"
	"fmt"
	"time"

	"go.bug.st/serial"
)

// Attempt to connect to usb serial
func ConnectUSBSerial() {
	// Error if no port selected
	if global.Ui.Ports.SelectedIndex() == -1 {
		dialogs.ShowError(errors.New("a port must be selected"))
		return
	}

	// Get baudrate from dropdown
	mode := &serial.Mode{
		BaudRate: global.Baudrates[global.Ui.Baudrate.SelectedIndex()],
		InitialStatusBits: &serial.ModemOutputBits{
			DTR: false,
			RTS: false,
		},
	}

	// Get port
	portName := global.Ui.Ports.Selected

	// Connect to serial port
	port, err := serial.Open(portName, mode)
	if err != nil {
		dialogs.ShowError(err)
		return
	}
	defer port.Close()

	// Re-assert to not reset on connection
	port.SetDTR(false)
	port.SetRTS(false)

	time.Sleep(2 * time.Second)

	message := "Hello World!\n"
	if n, err := port.Write([]byte(message)); err != nil {
		panic(err)
	} else {
		fmt.Printf("Sent %d bytes\n", n)
	}

	reader := bufio.NewReader(port)

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("RX:", line)
}
