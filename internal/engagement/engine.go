package engagement

import (
	"fmt"
	"ghost-hive/internal/geo"
	"ghost-hive/internal/models"
	"sort"
	"time"
)

type Engine struct {
	Hives []models.Hive
}

type HiveScore struct {
	Hive  models.Hive
	Score float64
}

func (e *Engine) PlanEngagement(threat models.Threat) (*models.Mission, error) {
	if len(e.Hives) == 0 {
		return nil, fmt.Errorf("no hives available")
	}

	weather := geo.GetCurrentWeather(threat.CurrentLocation)

	// Hard Thresholds for Safety
	if weather.WindSpeed > 100 {
		return nil, fmt.Errorf("ENVIRONMENTAL BLOCK: Hurricane force winds make interception impossible")
	}
	if weather.Temperature > 60 || weather.Temperature < -40 {
		return nil, fmt.Errorf("ENVIRONMENTAL BLOCK: Extreme temperature outside hardware operating range (%.1f C)", weather.Temperature)
	}
	if weather.Visibility < 10 && weather.Type == geo.WeatherSandstorm {
		return nil, fmt.Errorf("ENVIRONMENTAL BLOCK: Zero visibility in severe sandstorm")
	}

	var scores []HiveScore
	for _, h := range e.Hives {
		if h.Status != models.HiveStatusActive && h.Status != models.HiveStatusStandby {
			continue
		}

		effectiveCount := h.InterceptorsCount
		if h.InterceptorsCount <= 0 {
			continue
		}

		dist := geo.Distance(h.Location, threat.CurrentLocation)
		score := (1.0 / (dist + 1)) * 1000
		score += float64(effectiveCount) * 0.5

		// Environmental Penalties
		switch weather.Type {
		case geo.WeatherFog, geo.WeatherSandstorm:
			score *= 0.6 // Visual sensors degraded
		case geo.WeatherSnow:
			score *= 0.5 // Battery and aerodynamics degraded
		}

		if weather.Temperature > 45 {
			score *= 0.8 // Thermal throttling on electronics
		} else if weather.Temperature < -10 {
			score *= 0.7 // Battery chemistry efficiency drop
		}

		if weather.IsJamming {
			score *= 0.2
		}

		scores = append(scores, HiveScore{Hive: h, Score: score})
	}

	if len(scores) == 0 {
		return nil, fmt.Errorf("no suitable hives found for engagement under current environmental conditions")
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	bestHive := scores[0].Hive
	interceptionPt := geo.CalculateInterceptionPoint(bestHive.Location, threat, 200.0)

	mission := &models.Mission{
		ID:                fmt.Sprintf("MISS-%s-%s", bestHive.ID, threat.ID),
		ThreatID:          threat.ID,
		HiveIDs:           []string{bestHive.ID},
		Status:            models.MissionPlanned,
		StartTime:         time.Now(),
		InterceptionPt:    interceptionPt,
		ProjectedTargetPt: threat.ProjectedTarget,
		HiveSafeCrashSite: bestHive.SafeCrashSite,
	}

	return mission, nil
}
