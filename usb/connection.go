package usb

import (
	"bufio"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"go.bug.st/serial"
)

// Format of serial frame
type SerialFrame struct {
	Event    string         `json:"event"`
	Location string         `json:"location"`
	Payload  map[string]any `json:"payload"`
}

// Handles messaging and connection lifetime
type Connection struct {
	port   serial.Port
	reader *bufio.Reader

	mu sync.Mutex
}

// Create new connection object
func NewConnection(portName string, baud int) (*Connection, error) {
	// Setup baud and don't reset
	mode := &serial.Mode{
		BaudRate: baud,
		InitialStatusBits: &serial.ModemOutputBits{
			DTR: false,
			RTS: false,
		},
	}

	// Connect to serial port
	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, err
	}
	port.SetReadTimeout(50 * time.Millisecond)

	// Re-assert to not reset on connection
	port.SetDTR(false)
	port.SetRTS(false)

	// Create serial reader
	reader := bufio.NewReader(port)

	// Create connection object with port and reader
	connection := Connection{
		port:   port,
		reader: reader,
	}

	// Check if connection succeeded
	if _, err := connection.ping(); err != nil {
		connection.Disconnect() // Close serial port to prevent future "Port busy" errors
		return nil, err
	}

	return &connection, nil
}

// Disconnect connection
func (c *Connection) Disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.port != nil {
		c.port.Close()
		c.port = nil
	}
}

// Send message frame over serial and return response
func (c *Connection) Communicate(msg SerialFrame) (SerialFrame, error) {
	// Lock mutex so only one message in transit at time
	c.mu.Lock()
	defer c.mu.Unlock()

	// Send message
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
	if msg.Event == "error" {
		return SerialFrame{Event: "", Location: "", Payload: map[string]any{}}, errors.New(msg.Payload["status"].(string))
	}

	// fmt.Println(msg)

	return msg, nil
}

// Checks if the serial port is connected with ping messages
func (c *Connection) ping() (bool, error) {
	// Send ping message
	if _, err := c.Communicate(SerialFrame{Event: "get", Location: "ping", Payload: map[string]any{}}); err != nil {
		return false, err
	}

	return true, nil
}
