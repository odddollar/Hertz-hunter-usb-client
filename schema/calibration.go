package schema

import (
	"Hertz-Hunter-USB-Client/usb"
)

// Get calibrated values
func (s *Schema) GetCalibratedValues() (int, int, error) {
	frame := usb.SerialFrame{
		Event:    "get",
		Location: "calibration",
		Payload:  map[string]any{},
	}

	raw, err := s.connection.Communicate(frame)
	if err != nil {
		return 0, 0, err
	}

	return int(raw.Payload["low_rssi"].(float64)), int(raw.Payload["high_rssi"].(float64)), nil
}

// Set calibrated values
func (s *Schema) SetCalibratedValues(low, high int) error {
	frame := usb.SerialFrame{
		Event:    "post",
		Location: "calibration",
		Payload: map[string]any{
			"low_rssi":  low,
			"high_rssi": high,
		},
	}

	if _, err := s.connection.Communicate(frame); err != nil {
		return err
	}

	return nil
}
