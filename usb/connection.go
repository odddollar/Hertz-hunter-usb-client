// DISCLAIMER
//
// The connection/communication code here, particularly the Communicate() and receive()
// functions, operates on a "works" rather than "is correct" basis. It's hacky, but it's
// the best I could come up with.

package usb

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"go.bug.st/serial"
)

// Keep track of how long most recent receive() call took
var recentWait time.Duration

// Format of serial frame
type SerialFrame struct {
	Event    string         `json:"event"`
	Location string         `json:"location"`
	Payload  map[string]any `json:"payload"`
}

// Handles messaging and connection lifetime
type Connection struct {
	// Handle connection
	port       serial.Port
	maxRetries int

	// Handle connection logging
	logFile *os.File
	logger  *log.Logger

	// Handle preventing simultaneous messages
	mu sync.Mutex
}

// Create new connection object
func NewConnection(portName string, baud, maxRetries int) (*Connection, error) {
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

	// Create log file and logger
	logFile, err := os.OpenFile("usb_connection.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		port.Close()
		return nil, err
	}
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Create connection object with port
	connection := Connection{
		port:       port,
		maxRetries: maxRetries,
		logFile:    logFile,
		logger:     logger,
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

	// Purge and close connection stream
	c.port.ResetInputBuffer()
	c.port.ResetOutputBuffer()
	c.port.Close()

	// Close log file
	c.logFile.Close()
}

// Send message frame over serial and return response
func (c *Connection) Communicate(msg SerialFrame) (SerialFrame, error) {
	// Lock mutex so only one message in transit at time
	c.mu.Lock()
	defer c.mu.Unlock()

	var lastErr error

	for i := range c.maxRetries {
		// Drain serial to start buffer fresh
		c.port.ResetInputBuffer()
		c.port.ResetOutputBuffer()

		// Send message
		if err := c.send(msg); err != nil {
			lastErr = err
			c.logger.Printf("Retrying send (%d): %v", i, err)
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Read response
		rec, err := c.receive()
		if err != nil {
			lastErr = err
			c.logger.Printf("Retrying send/receive (%d, %s): %v", i, recentWait, err)
			time.Sleep(50 * time.Millisecond)
			continue
		}

		return rec, nil
	}

	c.logger.Printf("Communication failed after %d retries: %v", c.maxRetries, lastErr)
	return SerialFrame{}, lastErr
}

// Send message
func (c *Connection) send(msg SerialFrame) error {
	// Marshall struct
	data, err := json.Marshal(msg)
	if err != nil {
		c.logger.Printf("JSON marshalling error: %v", err)
		return err
	}

	// Ensure ends with newline
	data = append(data, '\n')

	// Write data to port
	_, err = c.port.Write(data)
	if err != nil {
		c.logger.Printf("Port write error: %v", err)
		return err
	}

	return nil
}

// Receive message
func (c *Connection) receive() (SerialFrame, error) {
	// Set overall message deadline
	// Port timeout used for single read() calls
	deadline := time.Now().Add(500 * time.Millisecond)

	// Keep track of how long this receive() call took
	startTime := time.Now()
	defer func() {
		recentWait = time.Since(startTime)
	}()

	// Buffers for reading data from serial
	var buffer bytes.Buffer
	tmp := make([]byte, 1024)

	readStarted := false

	// Retry receive message until deadline
	for time.Now().Before(deadline) {
		// Read bytes to tmp buffer
		nBytes, err := c.port.Read(tmp)
		if err != nil {
			c.logger.Printf("Port read error: %v", err)
			return SerialFrame{}, err
		}
		if nBytes == 0 {
			continue
		}

		// Process each received byte
		for i := range nBytes {
			b := tmp[i]

			// Wait for start of frame
			if !readStarted {
				if b == '{' {
					readStarted = true
					buffer.WriteByte(b)
				}
				continue
			}

			// Process full frame when newline received
			if b == '\n' {
				var msg SerialFrame
				if err := json.Unmarshal(buffer.Bytes(), &msg); err != nil {
					c.logger.Printf("JSON unmarshalling error: %v | raw: %s", err, buffer.String())
					return SerialFrame{}, err
				}
				if msg.Event == "error" {
					// Safe type assertion
					if status, ok := msg.Payload["status"].(string); ok {
						c.logger.Printf("Device error: %s", status)
						return SerialFrame{}, errors.New(status)
					}
					c.logger.Println("Unknown device error")
					return SerialFrame{}, errors.New("unknown device error")
				}

				return msg, nil
			}

			// Accumulate bytes
			buffer.WriteByte(b)
		}
	}

	c.logger.Println("Serial receive timeout")
	return SerialFrame{}, errors.New("serial receive timeout")
}

// Checks if the serial port is connected with ping messages
func (c *Connection) ping() (bool, error) {
	// Send ping message
	if _, err := c.Communicate(SerialFrame{Event: "get", Location: "ping", Payload: map[string]any{}}); err != nil {
		return false, err
	}

	return true, nil
}
