package models

import (
	"time"
)

type HiveStatus string

const (
	HiveStatusIdle        HiveStatus = "IDLE"
	HiveStatusStandby     HiveStatus = "STANDBY"
	HiveStatusActive      HiveStatus = "ACTIVE"
	HiveStatusInProduction HiveStatus = "IN_PRODUCTION"
	HiveStatusInTransport  HiveStatus = "IN_TRANSPORT"
	HiveStatusMaintenance  HiveStatus = "MAINTENANCE"
)

type ConnectivitySource string

const (
	ConnStarlink ConnectivitySource = "STARLINK"
	Conn4G       ConnectivitySource = "4G"
	ConnEthernet ConnectivitySource = "ETHERNET"
)

type HiveEnvironment struct {
	Type          string  `json:"type"` // ARCTIC, DESERT, AMPHIBIOUS, STANDARD
	Temperature   float64 `json:"temperature"`
	Pressure      float64 `json:"pressure"`
	SealIntegrity float64 `json:"seal_integrity"`
	HeaterLoad    float64 `json:"heater_load"`
}

type Hive struct {
	ID                 string             `json:"id"`
	Location           Coordinate         `json:"location"`
	SafeCrashSite      Coordinate         `json:"safe_crash_site"`
	Status             HiveStatus         `json:"status"`
	EnergyLevel        float64            `json:"energy_level"`
	EnergyDuration     time.Duration      `json:"energy_duration"`
	IsPluggedIn        bool               `json:"is_plugged_in"`
	Connectivity       map[ConnectivitySource]float64 `json:"connectivity"`
	Environment        HiveEnvironment    `json:"environment"`
	InterceptorsCount  int                `json:"interceptors_count"`
	StorageCount       int                `json:"storage_count"`
	Interceptors       []Interceptor      `json:"interceptors"`
	Magazines          []Magazine         `json:"magazines"`
	Breaches           []string           `json:"breaches"`
	Malfunctions       []string           `json:"malfunctions"`
	EstimatedReadyTime *time.Time         `json:"estimated_ready_time,omitempty"`
	LiveTrackingCoord  *Coordinate        `json:"live_tracking_coord,omitempty"`
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
	ThreatSingle   ThreatType = "SINGLE"
	ThreatSwarm    ThreatType = "SWARM"
	ThreatFriendly ThreatType = "FRIENDLY"
	ThreatCivilian ThreatType = "CIVILIAN"
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
	DetectedBy        string     `json:"detected_by"`
	TransponderCode   string     `json:"transponder_code,omitempty"`
}

type UserRole string

const (
	RoleOperator  UserRole = "OPERATOR"
	RoleCommander UserRole = "COMMANDER"
	RoleAdmin     UserRole = "ADMIN"
)

type User struct {
	ID   string   `json:"id"`
	Role UserRole `json:"role"`
}

type MissionStatus string

const (
	MissionPlanned   MissionStatus = "PLANNED"
	MissionActive    MissionStatus = "ACTIVE"
	MissionSuccess   MissionStatus = "SUCCESS"
	MissionFailed    MissionStatus = "FAILED"
	MissionIntercept MissionStatus = "INTERCEPTED"
)

type Mission struct {
	ID                string        `json:"id"`
	ThreatID          string        `json:"threat_id"`
	InterceptorIDs    []string      `json:"interceptor_ids"`
	HiveIDs           []string      `json:"hive_ids"`
	Status            MissionStatus `json:"status"`
	StartTime         time.Time     `json:"start_time"`
	InterceptionPt    Coordinate    `json:"interception_point"`
	ProjectedTargetPt Coordinate    `json:"projected_target_point"`
	HiveSafeCrashSite Coordinate    `json:"hive_safe_crash_site"`
	AuthorizedBy      string        `json:"authorized_by"`
}
