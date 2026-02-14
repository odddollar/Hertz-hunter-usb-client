package schema

import "Hertz-Hunter-USB-Client/usb"

// Set band to high or low band
func (s *Schema) SetBand(lowband bool) error {
	frame := usb.SerialFrame{
		Event:    "post",
		Location: "values",
		Payload: map[string]any{
			"lowband": lowband,
		},
	}

	if _, err := s.connection.Communicate(frame); err != nil {
		return err
	}

	return nil
}
