package usb

import (
	"Hertz-Hunter-USB-Client/global"
	"bufio"
	"encoding/json"
	"fmt"
	"sync"

	"go.bug.st/serial"
)

// Global connection struct for handling messaging and connection lifetime
type Connection struct {
	port   serial.Port
	reader *bufio.Reader

	mu sync.Mutex
}

// Disconnect connection
func (c *Connection) Disconnect() {
	global.EnableConnectionUI()
	c.port.Close()
}

// Send message
func (c *Connection) send(msg SerialFrame) error {
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
func (c *Connection) receive() (SerialFrame, error) {
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

// Sends message frame over serial and returns response
func (c *Connection) Communicate(msg SerialFrame) (SerialFrame, error) {
	// Lock mutex so only one message in transit at time
	c.mu.Lock()
	defer c.mu.Unlock()

	// Send ping message
	if err := c.send(msg); err != nil {
		return SerialFrame{Event: "", Location: "", Payload: map[string]any{}}, err
	}

	// Read response
	rec, err := c.receive()
	if err != nil {
		return SerialFrame{Event: "", Location: "", Payload: map[string]any{}}, err
	}

	return rec, nil
}

// Checks if the serial port is connected with ping messages
func (c *Connection) IsSerialConnected() (bool, error) {
	// Send ping message
	if _, err := c.Communicate(SerialFrame{Event: "get", Location: "ping", Payload: map[string]any{}}); err != nil {
		return false, err
	}

	return true, nil
}
