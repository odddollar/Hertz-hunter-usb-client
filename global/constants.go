package global

import (
	"Hertz-Hunter-USB-Client/schema"
	"time"
)

var (
	// Baudrates used in ui dropdown
	Baudrates       = []int{9600, 19200, 38400, 57600, 115200}
	DefaultBaudrate = 115200

	// Graph refresh intervals used in ui dropdown
	RefreshIntervals       = []time.Duration{100 * time.Millisecond, 250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second}
	DefaultRefreshInterval = 250 * time.Millisecond
)

// Dimensions for graph image
const (
	GraphWidth  = 998
	GraphHeight = 600
)

// Global schema object store
var Schema *schema.Schema
