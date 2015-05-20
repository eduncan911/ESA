package esa

// LibUsb interface is used to encapsulate the lower-level libusb implementation
type libUsb interface {
}

// defaultLibUsb implements the HidUsb interface for this package
type defaultLibUsb struct {
}

// newDefaultLibUsb instantiates a new DefaultLibUsb for the detected platform.
func newDefaultLibUsb() (libUsb, error) {
	return &defaultLibUsb{}, nil
}
