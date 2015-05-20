package main

import (
	"log"

	"github.com/eduncan911/esa"
)

var _esa esa.Esa

func main() {
	a := _esa.GetDevices()[0].Descriptor
	b := _esa.GetDevices()[0].Features[0].Descriptor
	c := _esa.GetDevices()[0].Features[0].Zones[0].Descriptor

	log.Println(a, b, c)
}

func init() {
	_esa = esa.New()
}
