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

	// 1. COLLATERAL ASSESSMENT: Decide Interceptor Type
	interceptionPt := geo.CalculateInterceptionPoint(models.Coordinate{}, threat, 200.0) // Mock base
	debris := geo.PredictDebrisFall(interceptionPt, models.Coordinate{Lat: 1, Lon: 0}) // Assume wind

	requiredType := models.TypeKinetic
	// If debris falls in high density area and threat is explosive-laden, use EXPLOSIVE interceptor for mid-air disposal
	if threat.PayloadDetected && !debris.IsSafe(0.8) {
		fmt.Println("ENGAGEMENT: High collateral risk detected. Escalating to EXPLOSIVE interceptor for mid-air disposal.")
		requiredType = models.TypeExplosive
	}

	var scores []HiveScore
	for _, h := range e.Hives {
		if h.Status != models.HiveStatusActive && h.Status != models.HiveStatusStandby {
			continue
		}

		// Check if Hive has the required type loaded
		hasType := false
		for _, m := range h.Magazines {
			if m.Type == requiredType && m.InterceptorsCount > 0 {
				hasType = true
				break
			}
		}
		if !hasType {
			continue
		}

		dist := geo.Distance(h.Location, threat.CurrentLocation)
		score := (1.0 / (dist + 1)) * 1000
		score += float64(h.InterceptorsCount) * 0.5

		if weather.WindSpeed > 20 { score *= 0.5 }
		if weather.IsJamming { score *= 0.2 }

		scores = append(scores, HiveScore{Hive: h, Score: score})
	}

	if len(scores) == 0 {
		return nil, fmt.Errorf("no suitable hives found with %s payload for this mission", requiredType)
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	bestHive := scores[0].Hive
	mission := &models.Mission{
		ID:                fmt.Sprintf("MISS-%s-%s", bestHive.ID, threat.ID),
		ThreatID:          threat.ID,
		InterceptorType:   requiredType,
		HiveIDs:           []string{bestHive.ID},
		Status:            models.MissionPlanned,
		StartTime:         time.Now(),
		InterceptionPt:    interceptionPt,
		ProjectedTargetPt: threat.ProjectedTarget,
		HiveSafeCrashSite: bestHive.SafeCrashSite,
	}

	return mission, nil
}
