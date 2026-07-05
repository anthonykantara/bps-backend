package nato

import (
	"fmt"
	"ghost-hive/internal/models"
)

// STANAG4586Message represents a simplified STANAG 4586 message structure
type STANAG4586Message struct {
	MsgID   int         `json:"msg_id"`
	Payload interface{} `json:"payload"`
}

const (
	MsgVehicleID            = 1
	MsgVehicleStatus        = 2
	MsgCUCSAuthorisation    = 3
	MsgLoiteringPosition    = 4
	MsgWaypoints            = 5
)

// TranslateToInternal converts a NATO message to our internal models
func TranslateToInternal(msg STANAG4586Message) (interface{}, error) {
	switch msg.MsgID {
	case MsgVehicleStatus:
		// Map STANAG vehicle status to Interceptor model
		return models.Interceptor{}, nil
	default:
		return nil, fmt.Errorf("unknown message ID: %d", msg.MsgID)
	}
}

// TranslateToNATO converts internal models to STANAG 4586 compatible format
func TranslateToNATO(obj interface{}) (STANAG4586Message, error) {
	switch v := obj.(type) {
	case models.Interceptor:
		return STANAG4586Message{
			MsgID: MsgVehicleStatus,
			Payload: map[string]interface{}{
				"vehicle_id": v.ID,
				"latitude":   v.Location.Lat,
				"longitude":  v.Location.Lon,
				"altitude":   v.Location.Alt,
				"status":     v.Status,
			},
		}, nil
	default:
		return STANAG4586Message{}, fmt.Errorf("unsupported type for NATO translation")
	}
}

// RadarMessage represents a NATO-compatible radar track (STANAG 4607/4609 simplified)
type RadarMessage struct {
	TrackID    string  `json:"track_id"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Alt        float64 `json:"alt"`
	Velocity   float64 `json:"velocity"`
	Heading    float64 `json:"heading"`
	TargetType string  `json:"target_type"`
}

func (r RadarMessage) ToThreat() models.Threat {
	return models.Threat{
		ID:              r.TrackID,
		DroneType:       r.TargetType,
		CurrentLocation: models.Coordinate{Lat: r.Lat, Lon: r.Lon, Alt: r.Alt},
		Speed:           r.Velocity,
		Altitude:        r.Alt,
	}
}
