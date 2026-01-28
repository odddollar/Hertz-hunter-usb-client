package usb

type SerialFrame struct {
	Event    string         `json:"event"`
	Location string         `json:"location"`
	Payload  map[string]any `json:"payload"`
}
