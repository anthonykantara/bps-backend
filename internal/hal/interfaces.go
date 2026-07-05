package hal

import (
	"context"
	"ghost-hive/internal/models"
)

type FlightController interface {
	Arm(ctx context.Context) error
	Disarm(ctx context.Context) error
	SetMode(ctx context.Context, mode string) error
	NavigateTo(ctx context.Context, target models.Coordinate) error
	GetTelemetry(ctx context.Context) (models.Coordinate, float64, error)
}

type MotorController interface {
	SetRPM(motorID int, rpm float64) error
	GetStatus() map[int]float64
}

type GimbalController interface {
	SetAngles(pitch, yaw, roll float64) error
	PointAt(ctx context.Context, target models.Coordinate) error
}

type SensorSuite interface {
	GetAcousticSignature() float64
	GetVisualDetection() []models.Threat
	GetEMSignature() float64
}
