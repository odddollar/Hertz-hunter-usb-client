package usb

import (
	"Hertz-Hunter-USB-Client/dialogs"
	"bufio"
	"encoding/json"
	"fmt"

	"go.bug.st/serial"
)

// Global connection struct for handling messaging and connection lifetime
type Connection struct {
	port   serial.Port
	reader *bufio.Reader
}

// Disconnect connection
func (c *Connection) Disconnect() {
	c.port.Close()
}

// Send message
func (c *Connection) Send(msg SerialFrame) error {
	// Marshall struct
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Ensure ends with newline
	data = append(data, '\n')

	// Write data to port
	_, err = c.port.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Receive message
func (c *Connection) Receive() (SerialFrame, error) {
	// Read serial until newline
	line, err := c.reader.ReadString('\n')
	if err != nil {
		return SerialFrame{Event: "", Location: "", Payload: map[string]any{}}, err
	}

	var msg SerialFrame

	// Unmarshal json to struct
	err = json.Unmarshal([]byte(line), &msg)
	if err != nil {
		return SerialFrame{Event: "", Location: "", Payload: map[string]any{}}, err
	}

	fmt.Println(msg)

	return msg, nil
}

// Checks if the serial port is connected with ping messages
func (c *Connection) IsSerialConnected() bool {
	// Send ping message
	if err := c.Send(SerialFrame{Event: "get", Location: "ping", Payload: map[string]any{}}); err != nil {
		dialogs.ShowError(err)
		return false
	}

	// Read response
	if _, err := c.Receive(); err != nil {
		dialogs.ShowError(err)
		return false
	}

	return true
}
