package schema

import (
	"Hertz-Hunter-USB-Client/usb"
	"context"
	"fmt"
	"time"
)

// Handles messaging device with required schema
type Schema struct {
	connection    *usb.Connection
	pollingCancel context.CancelFunc
}

// Create new schema object
func NewSchema(con *usb.Connection) *Schema {
	return &Schema{
		connection: con,
	}
}

// Start polling for rssi values from device
func (c *Schema) StartPollValues(period time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())
	c.pollingCancel = cancel

	go func() {
		// Start ticker
		ticker := time.NewTicker(period)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				_, err := c.connection.Communicate(usb.SerialFrame{Event: "get", Location: "values", Payload: map[string]any{}})
				if err != nil {
					fmt.Printf("Periodic communication error: %v\n", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Cancels rssi polling
func (c *Schema) StopPollValues() {
	if c.pollingCancel != nil {
		c.pollingCancel()
		c.pollingCancel = nil
	}
}
