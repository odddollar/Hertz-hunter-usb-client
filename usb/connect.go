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

// Globally store port and
var port serial.Port
var serialReader *bufio.Reader

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
	var err error
	port, err = serial.Open(portName, mode)
	if err != nil {
		dialogs.ShowError(err)
		return
	}
	defer port.Close()
	port.SetReadTimeout(50 * time.Millisecond)

	// Re-assert to not reset on connection
	port.SetDTR(false)
	port.SetRTS(false)

	// Create serial reader
	serialReader = bufio.NewReader(port)

	time.Sleep(2 * time.Second)

	isSerialConnected()

	// message := "Hello World!\n"
	// if n, err := port.Write([]byte(message)); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("Sent %d bytes\n", n)
	// }

	// reader := bufio.NewReader(port)

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("RX:", line)
}

// Checks if the serial port is connected with ping messages
func isSerialConnected() bool {
	// Send ping message
	message := `{"event":"get","location":"ping","payload":{}}`
	if _, err := port.Write([]byte(message + "\n")); err != nil {
		dialogs.ShowError(err)
		return false
	}

	// Read response
	if line, err := serialReader.ReadString('\n'); err != nil {
		dialogs.ShowError(err)
		return false
	} else {
		fmt.Println(line)
	}

	return true
}
