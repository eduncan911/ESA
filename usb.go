package esa

// Usb is the interface to all underlying USB functions
type usb interface {
}

// defaultUsb implements the Usb interface for this package
type defaultUsb struct {
	hidusb hidUsb
	libusb libUsb
}

// newDefaultUsb instantiates a new DefaultUsb
func newDefaultUsb(h hidUsb, l libUsb) (usb, error) {
	return &defaultUsb{
		hidusb: h,
		libusb: l,
	}, nil
}
