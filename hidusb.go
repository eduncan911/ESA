package esa

type hidElement struct {
	Descriptor  string
	Range       hidRange
	Position    hidXYZPosition
	SpatialZone string
	LEDColor    string
	Status      string
}

type hidRange struct {
	Minimal int
	Maximum int
	Warn    int
}

type hidXYZPosition struct {
	X int
	Y int
	Z int
}

// HidUsb interface is used to encapsulate common HID commands
type hidUsb interface {
}

// DefaultHidUsb implements the HidUsb interface for this package
type defaultHidUsb struct {
}

// newDefaultHidUsb instantiates a new DefaultHidUsb for the detected platform.
func newDefaultHidUsb() (hidUsb, error) {
	return &defaultHidUsb{}, nil
}
