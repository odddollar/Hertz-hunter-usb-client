package usb

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"bufio"
	"fmt"

	"go.bug.st/serial"
)

// Global connection struct for handling messaging and connection lifetime
type Connection struct {
	Port   serial.Port
	Reader *bufio.Reader
}

// Disconnect connection
func (c *Connection) Disconnect() {
	c.Port.Close()
}

// Checks if the serial port is connected with ping messages
func (c *Connection) IsSerialConnected() bool {
	// Send ping message
	message := `{"event":"get","location":"ping","payload":{}}`
	if _, err := c.Port.Write([]byte(message + "\n")); err != nil {
		dialogs.ShowError(err)
		return false
	}

	// Read response
	if line, err := c.Reader.ReadString('\n'); err != nil {
		dialogs.ShowError(err)
		return false
	} else {
		fmt.Println(line)
	}

	return true
}
