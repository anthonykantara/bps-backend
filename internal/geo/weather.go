package geo

import "ghost-hive/internal/models"

type WeatherCondition struct {
	WindSpeed float64
	Visibility float64
	IsJamming bool
}

func GetCurrentWeather(loc models.Coordinate) WeatherCondition {
	// Placeholder for real-time weather API integration
	return WeatherCondition{
		WindSpeed: 5.0,
		Visibility: 10000.0,
		IsJamming: false,
	}
}
