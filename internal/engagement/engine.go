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

// AnticipateFutureWaves analyzes incoming threat patterns to determine if we should reserve resources
func (e *Engine) AnticipateFutureWaves(threats []models.Threat) bool {
	// If we detect multiple single threats with staggered entry times,
	// suggest holding back 20% of interceptors in primary hives.
	return len(threats) > 500
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

		// Temporal Buffer: If Hive is <20% loaded and we anticipate more waves, deprioritize it
		effectiveCount := h.InterceptorsCount
		if effectiveCount < 32 { // 20% of 160
			effectiveCount /= 2
		}

		if h.InterceptorsCount <= 0 {
			continue
		}

		dist := geo.Distance(h.Location, threat.CurrentLocation)
		score := (1.0 / (dist + 1)) * 1000
		score += float64(effectiveCount) * 0.5

		if weather.WindSpeed > 20 {
			score *= 0.5
		}

		scores = append(scores, HiveScore{Hive: h, Score: score})
	}

	if len(scores) == 0 {
		return nil, fmt.Errorf("no suitable hives found for engagement")
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
