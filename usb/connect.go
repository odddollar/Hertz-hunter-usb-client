package usb

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"Hertz-Hunter-USB-Client/global"
	"bufio"
	"errors"
	"time"

	"go.bug.st/serial"
)

// Global connection object store
var connection Connection

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
	port, err := serial.Open(portName, mode)
	if err != nil {
		dialogs.ShowError(err)
		return
	}
	port.SetReadTimeout(50 * time.Millisecond)

	// Re-assert to not reset on connection
	port.SetDTR(false)
	port.SetRTS(false)

	// Create serial reader
	reader := bufio.NewReader(port)

	// Create connection object with port and reader
	connection = Connection{
		Port:   port,
		Reader: reader,
	}

	// Show success message if connected
	if connection.IsSerialConnected() {
		dialogs.ShowSuccess("Successfully connected to port")
	} else {
		dialogs.ShowError(errors.New("failed to connect to port"))
	}
}
