package data

import (
	"fmt"
	"time"
)

func StartExternalDatasetIngestor() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			fmt.Println("Ingestor: Pulling latest drone specs from OpenSky/Manufacturer datasets...")
			// Simulate updating GlobalDroneDB with new "Threat Silhouettes"
			GlobalDroneDB["NEW-THREAT-X"] = DroneSpecs{
				Model:    "Stealth-X",
				MaxSpeed: 80.0,
				RCS:      0.01,
				AltRange: [2]float64{10, 1000},
			}
		}
	}()
}
