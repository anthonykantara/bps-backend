package geo

import (
	"fmt"
	"ghost-hive/internal/models"
)

type DebrisField struct {
	Center   models.Coordinate
	RadiusM  float64
	Severity float64
}

func PredictDebrisFall(impact models.Coordinate, wind models.Coordinate) DebrisField {
	// Simple ballistics projection based on impact altitude and wind vector
	// 100m fall per 1m/s wind drift (placeholder)
	driftLat := wind.Lat * 0.001
	driftLon := wind.Lon * 0.001

	return DebrisField{
		Center: models.Coordinate{
			Lat: impact.Lat + driftLat,
			Lon: impact.Lon + driftLon,
			Alt: 0,
		},
		RadiusM: impact.Alt * 0.1, // Expansion based on altitude
		Severity: 0.8,
	}
}

func (d *DebrisField) IsSafe(populationDensity float64) bool {
	if populationDensity > 0.5 && d.Severity > 0.5 {
		fmt.Println("COLLATERAL WARNING: High debris risk in populated sector.")
		return false
	}
	return true
}
