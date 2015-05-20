package esa

// Device represents the physical component that implements the ESA standard.
type Device struct {
	Descriptor string
	Features   []Feature
}

type Feature struct {
	Descriptor  string
	FeatureType FeatureType
	ReportID    int8
	Zones       []Zone
}

// FeatureType defines what type of Device is configured. (i.e. Chassis, PowerSupply, WaterCooler)
type FeatureType int8

const (
	// Chassis is a feature type that must implements a minimal of:
	//
	Chassis FeatureType = iota
	// PowerSupply is a feature type that implements a minimal of:
	//
	PowerSupply
	// WaterCooler is a feature type that implements a minimal of:
	// (1) thermal zone
	WaterCooler
)

type Zone struct {
	Descriptor string
	ReportID   int8
	ZoneType   ZoneType
}

type ZoneType int8

const (
	Controller ZoneType = iota
	Sensor
)
