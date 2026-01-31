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
	c.connection.Disconnect()

	if c.pollingCancel != nil {
		c.pollingCancel()
		c.pollingCancel = nil
	}
}

// Start polling for rssi values from device
func (c *Schema) StartPollValues(period time.Duration) <-chan error {
	ctx, cancel := context.WithCancel(context.Background())
	c.pollingCancel = cancel

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		// Start ticker
		ticker := time.NewTicker(period)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				_, err := c.connection.Communicate(usb.SerialFrame{Event: "get", Location: "values", Payload: map[string]any{}})
				if err != nil {
					errCh <- err
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return errCh
}
