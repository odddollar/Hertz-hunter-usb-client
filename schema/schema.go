package schema

import (
	"Hertz-Hunter-USB-Client/usb"
	"context"
	"time"
)

// Handles messaging device with required schema
type Schema struct {
	connection    *usb.Connection
	pollingCancel context.CancelFunc
}

// Create new schema object
func NewSchema(portName string, baud int) (*Schema, error) {
	con, err := usb.NewConnection(portName, baud)
	if err != nil {
		return nil, err
	}

	return &Schema{
		connection: con,
	}, nil
}

// Cancels rssi polling and disconnects device
func (c *Schema) Stop() {
	if c.pollingCancel != nil {
		c.pollingCancel()
		c.pollingCancel = nil
	}

	if c.connection != nil {
		c.connection.Disconnect()
		c.connection = nil
	}
}

// Start polling for rssi values from device
func (c *Schema) StartPollValues(period time.Duration) (<-chan []int, <-chan error) {
	ctx, cancel := context.WithCancel(context.Background())
	c.pollingCancel = cancel

	valuesCh := make(chan []int)
	errCh := make(chan error, 1)

	go func() {
		defer close(valuesCh)
		defer close(errCh)

		// Start ticker
		ticker := time.NewTicker(period)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Get values data
				data, err := c.connection.Communicate(usb.SerialFrame{Event: "get", Location: "values", Payload: map[string]any{}})
				if err != nil {
					errCh <- err
					return
				}

				// Send data back over channel
				values, _ := data.Payload["values"].([]int)
				valuesCh <- values
			case <-ctx.Done():
				return
			}
		}
	}()

	return valuesCh, errCh
}
