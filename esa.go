package esa

type Esa interface {
	GetDevices() *[]Device
}

// New constructs a concrete instance of the Esa interface.  This is safe to call even
// with no Esa devices connected (GetDevices() will just return 0 devices).
//
// New is only allowed to be called once per execution.  Subsequent calls will
// will be skipped and return the same instance.
//
// Following best practices for IoC and unit testing, this should only be
// called from within your init() method of your package or executable and set
// to a global instance.  Then you can mock the Esa interface throughout your tests
// since Go tests do not execute init().
//
//  // global instance of the mockable Esa interface
//  var _esa esa.Esa
//
//  func init() {
//    esa = esa.New()
//  }
//
//  func main() {
//    devices := esa.GetDevices()
//    for _, d := range devices {
//      fmt.Println(d.Vendor)
//    }
//  }
func New() Esa {

	iesa := &defaultEsa{}

	// create an instance of the hidusb interface for this platform
	h, err := newDefaultHidUsb()
	if err != nil {
		return esa
	}

	// create an instance of the libusb interface for this platform
	l, err := newDefaultLibUsb()
	if err != nil {
		return esa
	}

	// new up the usb defaults
	usb, err := newDefaultUsb(h, l)
	if err != nil {
		return esa
	}

	// set the dependencies
	iesa.usb = usb

	// set the global instance
	esa = iesa

	return esa
}

// GetDevices returns a new collection of detected ESA devices.
func GetDevices() []Device {
	return *esa.GetDevices()
}

type defaultEsa struct {
	Devices *[]Device

	usb usb
}

func (e *defaultEsa) GetDevices() *[]Device {
	return e.Devices
}

var esa Esa
