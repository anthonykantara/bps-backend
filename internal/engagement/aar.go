package engagement

import (
	"fmt"
	"ghost-hive/internal/models"
	"time"
)

type AfterActionReport struct {
	MissionID   string
	StartTime   time.Time
	EndTime     time.Time
	ThreatLevel string
	Outcome     string
	Efficiency  float64 // Interceptors per threat
}

func GenerateAAR(mission models.Mission) AfterActionReport {
	return AfterActionReport{
		MissionID:   mission.ID,
		StartTime:   mission.StartTime,
		EndTime:     time.Now(),
		ThreatLevel: "HIGH",
		Outcome:     "SUCCESS - Threat Destroyed",
		Efficiency:  float64(len(mission.InterceptorIDs)),
	}
}

func (a AfterActionReport) Summary() string {
	return fmt.Sprintf("Mission %s: %s at %s", a.MissionID, a.Outcome, a.EndTime.Format(time.RFC3339))
}
