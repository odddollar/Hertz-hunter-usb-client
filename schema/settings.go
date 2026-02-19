package schema

import "Hertz-Hunter-USB-Client/usb"

// Get calibrated values
// Returns scan_interval_index, buzzer_index, battery_alarm_index
func (s *Schema) GetSettingsIndices() (int, int, int, error) {
	frame := usb.SerialFrame{
		Event:    "get",
		Location: "settings",
		Payload:  map[string]any{},
	}

	raw, err := s.connection.Communicate(frame)
	if err != nil {
		return 0, 0, 0, err
	}

	return int(raw.Payload["scan_interval_index"].(float64)),
		int(raw.Payload["buzzer_index"].(float64)),
		int(raw.Payload["battery_alarm_index"].(float64)),
		nil
}
