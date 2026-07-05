package geo

import (
	"ghost-hive/internal/models"
)

type KalmanFilter struct {
	State        models.Coordinate // Estimated [Lat, Lon, Alt]
	Covariance   float64           // Uncertainty
	ProcessNoise float64
	SensorNoise  float64
}

func NewKalmanFilter(initial models.Coordinate) *KalmanFilter {
	return &KalmanFilter{
		State:        initial,
		Covariance:   1.0,
		ProcessNoise: 0.1,
		SensorNoise:  0.5,
	}
}

// Update merges a noisy measurement into the state estimate
func (k *KalmanFilter) Update(measurement models.Coordinate) models.Coordinate {
	// 1. Prediction Step
	k.Covariance += k.ProcessNoise

	// 2. Innovation / Measurement Update
	kalmanGain := k.Covariance / (k.Covariance + k.SensorNoise)

	k.State.Lat += kalmanGain * (measurement.Lat - k.State.Lat)
	k.State.Lon += kalmanGain * (measurement.Lon - k.State.Lon)
	k.State.Alt += kalmanGain * (measurement.Alt - k.State.Alt)

	k.Covariance *= (1 - kalmanGain)

	return k.State
}

// MergeSensors handles data from GPS, IMU, and Visual sources
func MergeSensors(gps, imu, visual models.Coordinate) models.Coordinate {
	// Weighting based on sensor reliability (Placeholders)
	return models.Coordinate{
		Lat: (gps.Lat*0.4 + imu.Lat*0.2 + visual.Lat*0.4),
		Lon: (gps.Lon*0.4 + imu.Lon*0.2 + visual.Lon*0.4),
		Alt: (gps.Alt*0.3 + imu.Alt*0.2 + visual.Alt*0.5),
	}
}
