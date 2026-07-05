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

	var scores []HiveScore
	for _, h := range e.Hives {
		if h.Status != models.HiveStatusActive && h.Status != models.HiveStatusStandby {
			continue
		}
		if h.InterceptorsCount <= 0 {
			continue
		}

		dist := geo.Distance(h.Location, threat.CurrentLocation)
		score := (1.0 / (dist + 1)) * 1000
		score += float64(h.InterceptorsCount) * 0.5

		// Weather Impact: Penalty for high wind or low visibility
		if weather.WindSpeed > 20 {
			score *= 0.5 // 50% penalty for high wind
		}
		if weather.IsJamming {
			score *= 0.2 // Heavy penalty for jamming zones
		}

		scores = append(scores, HiveScore{Hive: h, Score: score})
	}

	if len(scores) == 0 {
		return nil, fmt.Errorf("no suitable hives found for engagement (weather/jamming too severe)")
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	bestHive := scores[0].Hive

	// Interception point adjusted for wind
	interceptionPt := geo.CalculateInterceptionPoint(bestHive.Location, threat, 200.0)
	if weather.WindSpeed > 10 {
		interceptionPt.Lat += 0.0001 // Drift compensation
	}

	mission := &models.Mission{
		ID:                fmt.Sprintf("MISS-%s-%s", bestHive.ID, threat.ID),
		ThreatID:          threat.ID,
		HiveIDs:           []string{bestHive.ID},
		InterceptorIDs:    []string{bestHive.Interceptors[0].ID},
		Status:            models.MissionPlanned,
		StartTime:         time.Now(),
		InterceptionPt:    interceptionPt,
		ProjectedTargetPt: threat.ProjectedTarget,
		HiveSafeCrashSite: bestHive.SafeCrashSite,
	}

	return mission, nil
}
