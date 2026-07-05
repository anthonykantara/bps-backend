package hive

import (
	"fmt"
	"ghost-hive/internal/models"
)

type DecoySystem struct {
	Decoys []DecoyHive
}

func (s *DecoySystem) ActivateCountermeasures(threatOrigin models.Coordinate) {
	fmt.Println("C2: Enemy reconnaissance detected. Activating Decoy Hive network...")
	for _, d := range s.Decoys {
		fmt.Printf("Decoy %s: Emitting signature spoofing data for Hive location %+v\n", d.ID, d.Location)
	}
}
