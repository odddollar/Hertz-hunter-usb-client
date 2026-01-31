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

// Used to return lowband state with values on polling
type ValuesResult struct {
	Values  []int
	Lowband bool
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
func (s *Schema) Stop() {
	if s.pollingCancel != nil {
		s.pollingCancel()
		s.pollingCancel = nil
	}

	if s.connection != nil {
		s.connection.Disconnect()
		s.connection = nil
	}
}

// Start polling for rssi values from device
func (s *Schema) StartPollValues(period time.Duration) (<-chan ValuesResult, <-chan error) {
	ctx, cancel := context.WithCancel(context.Background())
	s.pollingCancel = cancel

	valuesCh := make(chan ValuesResult)
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
				data, err := s.connection.Communicate(usb.SerialFrame{Event: "get", Location: "values", Payload: map[string]any{}})
				if err != nil {
					errCh <- err
					return
				}

				// Convert data to proper type
				raw, _ := data.Payload["values"].([]any)
				values := make([]int, len(raw))
				for i, v := range raw {
					f, _ := v.(float64)
					values[i] = int(f)
				}

				// Get lowband state
				lowband, _ := data.Payload["lowband"].(bool)

				// Send data over channel
				valuesCh <- ValuesResult{
					Values:  values,
					Lowband: lowband,
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return valuesCh, errCh
}

// Get calibrated values
func (s *Schema) GetCalibratedValues() (int, int, error) {
	raw, err := s.connection.Communicate(usb.SerialFrame{Event: "get", Location: "calibration", Payload: map[string]any{}})
	if err != nil {
		return 0, 0, err
	}

	return int(raw.Payload["low_rssi"].(float64)), int(raw.Payload["high_rssi"].(float64)), nil
}
