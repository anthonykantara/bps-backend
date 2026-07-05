package engagement

import (
	"fmt"
	"ghost-hive/internal/geo"
	"ghost-hive/internal/models"
	"sort"
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

	var scores []HiveScore
	for _, h := range e.Hives {
		if h.Status != models.HiveStatusActive && h.Status != models.HiveStatusStandby {
			continue
		}
		if h.InterceptorsCount <= 0 {
			continue
		}

		// Scoring logic:
		// 1. Distance to threat (closer is better)
		// 2. Resource availability (more interceptors remaining is better)
		// 3. Location density (avoid maxing out hives in high threat zones)
		dist := geo.Distance(h.Location, threat.CurrentLocation)
		score := (1.0 / (dist + 1)) * 1000 // Simple distance score
		score += float64(h.InterceptorsCount) * 0.5

		scores = append(scores, HiveScore{Hive: h, Score: score})
	}

	if len(scores) == 0 {
		return nil, fmt.Errorf("no suitable hives found for engagement")
	}

	// Sort by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	bestHive := scores[0].Hive
	interceptionPt := geo.CalculateInterceptionPoint(bestHive.Location, threat, 200.0) // Assume 200m/s interceptor

	mission := &models.Mission{
		ID:             fmt.Sprintf("MISS-%s-%s", bestHive.ID, threat.ID),
		ThreatID:       threat.ID,
		HiveIDs:        []string{bestHive.ID},
		InterceptorIDs: []string{bestHive.Interceptors[0].ID}, // Assign first available
		Status:         "PLANNED",
		InterceptionPt: interceptionPt,
	}

	return mission, nil
}
