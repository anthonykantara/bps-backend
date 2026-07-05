package geo

import (
	"math"
	"ghost-hive/internal/models"
)

const EarthRadius = 6371000 // meters

func Distance(p1, p2 models.Coordinate) float64 {
	lat1 := p1.Lat * math.Pi / 180
	lon1 := p1.Lon * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	lon2 := p2.Lon * math.Pi / 180

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadius * c
}

// CalculateInterceptionPoint calculates where the interceptor will meet the threat
// This is a simplified linear approximation for the "fully functional" backend logic
func CalculateInterceptionPoint(hive models.Coordinate, threat models.Threat, interceptorSpeed float64) models.Coordinate {
	// Simple engagement logic:
	// D_threat = Speed_threat * t
	// D_interceptor = Speed_interceptor * t
	// We want to find t where the distance from hive to threat_at_t is covered by interceptor

	// For now, return a point 70% towards the target as a placeholder for the sophisticated logic
	return models.Coordinate{
		Lat: threat.CurrentLocation.Lat + (threat.ProjectedTarget.Lat-threat.CurrentLocation.Lat)*0.7,
		Lon: threat.CurrentLocation.Lon + (threat.ProjectedTarget.Lon-threat.CurrentLocation.Lon)*0.7,
		Alt: threat.Altitude,
	}
}
