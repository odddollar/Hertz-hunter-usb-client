package schema

import "Hertz-Hunter-USB-Client/usb"

// Get settings indices
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

// Set settings indices
func (s *Schema) SetSettingsIndices(scan_interval_index, buzzer_index, battery_alarm_index int) error {
	frame := usb.SerialFrame{
		Event:    "post",
		Location: "settings",
		Payload: map[string]any{
			"scan_interval_index": scan_interval_index,
			"buzzer_index":        buzzer_index,
			"battery_alarm_index": battery_alarm_index,
		},
	}

	if _, err := s.connection.Communicate(frame); err != nil {
		return err
	}

	return nil
}
