package data

type DroneSpecs struct {
	Model       string
	MaxSpeed    float64 // m/s
	RCS         float64 // Radar Cross Section
	AltRange    [2]float64
	VulnerabilityPoints []string
}

var GlobalDroneDB = map[string]DroneSpecs{
	"SHARED-136": {
		Model:    "Shahed 136",
		MaxSpeed: 50.0,
		RCS:      0.5,
		AltRange: [2]float64{100, 4000},
		VulnerabilityPoints: []string{"Engine", "Payload", "Control Surface"},
	},
	"ORLAN-10": {
		Model:    "Orlan-10",
		MaxSpeed: 41.0,
		RCS:      0.2,
		AltRange: [2]float64{500, 5000},
		VulnerabilityPoints: []string{"Camera Gimbal", "Fuel Tank"},
	},
}

func GetDroneSpecs(model string) *DroneSpecs {
	if spec, ok := GlobalDroneDB[model]; ok {
		return &spec
	}
	return nil
}
