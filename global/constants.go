package global

import "time"

var (
	Baudrates       = []int{9600, 19200, 38400, 57600, 115200}
	DefaultBaudrate = 115200

	RefreshIntervals       = []time.Duration{250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second}
	DefaultRefreshInterval = 500 * time.Millisecond
)

const (
	GraphWidth  = 1000
	GraphHeight = 600
)
