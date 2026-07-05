package interceptor

import "fmt"

type PayloadType string

const (
	PayloadKinetic   PayloadType = "KINETIC"
	PayloadElectronic PayloadType = "EW_JAMMER"
	PayloadSensor     PayloadType = "SENS_RELAY"
)

type PayloadManager struct {
	Type PayloadType
}

func (p *PayloadManager) Execute() {
	switch p.Type {
	case PayloadElectronic:
		fmt.Println("EW Payload: Emitting broadband noise to disrupt threat command link...")
	case PayloadKinetic:
		fmt.Println("Kinetic Payload: Preparing for terminal impact.")
	}
}
