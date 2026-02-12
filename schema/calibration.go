package schema

import "Hertz-Hunter-USB-Client/usb"

// Get calibrated values
func (s *Schema) GetCalibratedValues() (int, int, error) {
	raw, err := s.connection.Communicate(usb.SerialFrame{Event: "get", Location: "calibration", Payload: map[string]any{}})
	if err != nil {
		return 0, 0, err
	}

	return int(raw.Payload["low_rssi"].(float64)), int(raw.Payload["high_rssi"].(float64)), nil
}
