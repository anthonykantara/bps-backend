package hive

import "ghost-hive/internal/models"

type DecoyHive struct {
	models.Hive
	IsEmissionOnly bool
}

func NewDecoy(id string, loc models.Coordinate) DecoyHive {
	return DecoyHive{
		Hive: models.Hive{
			ID: id,
			Location: loc,
			Status: models.HiveStatusActive,
			Connectivity: map[models.ConnectivitySource]float64{
				models.ConnStarlink: 90.0,
			},
		},
		IsEmissionOnly: true,
	}
}
