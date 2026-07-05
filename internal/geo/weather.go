package geo

import (
	"ghost-hive/internal/models"
	"sync"
)

type WeatherType string

const (
	WeatherClear     WeatherType = "CLEAR"
	WeatherFog       WeatherType = "FOG"
	WeatherSandstorm WeatherType = "SANDSTORM"
	WeatherSnow      WeatherType = "SNOW"
	WeatherHurricane WeatherType = "HURRICANE"
)

type WeatherCondition struct {
	Type        WeatherType
	Name        string
	Temperature float64
	WindSpeed   float64
	Visibility  float64
	IsJamming   bool
}

var (
	currentWeather = WeatherCondition{
		Type:        WeatherClear,
		Name:        "DEFAULT CLEAR",
		Temperature: 20.0,
		WindSpeed:   5.0,
		Visibility:  10000.0,
		IsJamming:   false,
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
