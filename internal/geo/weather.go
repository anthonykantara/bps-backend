package geo

import (
	"ghost-hive/internal/models"
	"sync"
)

type WeatherCondition struct {
	Name       string
	WindSpeed  float64
	Visibility float64
	IsJamming  bool
}

var (
	currentWeather = WeatherCondition{
		Name:       "CLEAR",
		WindSpeed:  5.0,
		Visibility: 10000.0,
		IsJamming:  false,
	}
	weatherMu sync.RWMutex
)

func SetWeather(w WeatherCondition) {
	weatherMu.Lock()
	defer weatherMu.Unlock()
	currentWeather = w
}

func GetCurrentWeather(loc models.Coordinate) WeatherCondition {
	weatherMu.RLock()
	defer weatherMu.RUnlock()
	return currentWeather
}
