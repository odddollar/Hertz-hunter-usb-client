package global

import (
	"Hertz-Hunter-USB-Client/schema"
	"Hertz-Hunter-USB-Client/usb"
	"time"
)

var (
	// Baudrates used in ui dropdown
	Baudrates       = []int{9600, 19200, 38400, 57600, 115200}
	DefaultBaudrate = 115200

	// Graph refresh intervals used in ui dropdown
	RefreshIntervals       = []time.Duration{250 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second}
	DefaultRefreshInterval = 500 * time.Millisecond
)

// Dimensions for graph image
const (
	GraphWidth  = 1000
	GraphHeight = 600
)

// Global connection objects store
var Connection *usb.Connection
var Schema *schema.Schema
