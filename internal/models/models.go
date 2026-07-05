package models

import (
	"time"
)

type HiveStatus string

const (
	HiveStatusIdle      HiveStatus = "IDLE"
	HiveStatusStandby   HiveStatus = "STANDBY"
	HiveStatusActive    HiveStatus = "ACTIVE"
	HiveStatusTransport HiveStatus = "TRANSPORT"
	HiveStatusMaintenance HiveStatus = "MAINTENANCE"
)

type ConnectivitySource string

const (
	ConnStarlink ConnectivitySource = "STARLINK"
	Conn4G       ConnectivitySource = "4G"
	ConnEthernet ConnectivitySource = "ETHERNET"
)

type Hive struct {
	ID                 string             `json:"id"`
	Location           Coordinate         `json:"location"`
	Status             HiveStatus         `json:"status"`
	EnergyLevel        float64            `json:"energy_level"` // Fuel %
	EnergyDuration     time.Duration      `json:"energy_duration"`
	IsPluggedIn        bool               `json:"is_plugged_in"`
	Connectivity       map[ConnectivitySource]float64 `json:"connectivity"` // Source to signal strength
	Temperature        float64            `json:"temperature"`
	InterceptorsCount  int                `json:"interceptors_count"`
	Interceptors       []Interceptor      `json:"interceptors"`
	Magazines          []Magazine         `json:"magazines"`
	Breaches           []string           `json:"breaches"`
	Malfunctions       []string           `json:"malfunctions"`
	EstimatedReadyTime *time.Time         `json:"estimated_ready_time,omitempty"`
}

type Magazine struct {
	ID                string `json:"id"`
	InterceptorsCount int    `json:"interceptors_count"`
}

type InterceptorStatus string

const (
	InterceptorStatusIdle      InterceptorStatus = "IDLE"
	InterceptorStatusCharging  InterceptorStatus = "CHARGING"
	InterceptorStatusReady     InterceptorStatus = "READY"
	InterceptorStatusLaunched  InterceptorStatus = "LAUNCHED"
	InterceptorStatusEngaged   InterceptorStatus = "ENGAGED"
	InterceptorStatusDestroyed InterceptorStatus = "DESTROYED"
	InterceptorStatusFailed    InterceptorStatus = "FAILED"
)

type Interceptor struct {
	ID           string            `json:"id"`
	HiveID       string            `json:"hive_id"`
	Status       InterceptorStatus `json:"status"`
	BatteryLevel float64           `json:"battery_level"`
	Health       float64           `json:"health"`
	Location     Coordinate        `json:"location"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Alt float64 `json:"alt"`
}

type ThreatType string

const (
	ThreatSingle ThreatType = "SINGLE"
	ThreatSwarm  ThreatType = "SWARM"
)

type Threat struct {
	ID                string     `json:"id"`
	Type              ThreatType `json:"type"`
	DroneType         string     `json:"drone_type"`
	Origin            Coordinate `json:"origin"`
	CurrentLocation   Coordinate `json:"current_location"`
	ProjectedTarget   Coordinate `json:"projected_target"`
	Speed             float64    `json:"speed"`
	Altitude          float64    `json:"altitude"`
	LastUpdated       time.Time  `json:"last_updated"`
	DetectedBy        string     `json:"detected_by"` // Radar ID
}

type Mission struct {
	ID             string     `json:"id"`
	ThreatID       string     `json:"threat_id"`
	InterceptorIDs []string   `json:"interceptor_ids"`
	HiveIDs        []string   `json:"hive_ids"`
	Status         string     `json:"status"`
	StartTime      time.Time  `json:"start_time"`
	InterceptionPt Coordinate `json:"interception_point"`
}
