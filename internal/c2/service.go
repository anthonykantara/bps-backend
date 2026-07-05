package c2

import (
	"context"
	"fmt"
	"ghost-hive/internal/engagement"
	"ghost-hive/internal/models"
	"sync"
)

type C2Server struct {
	mu          sync.RWMutex
	Hives       map[string]models.Hive
	Threats     map[string]models.Threat
	Missions    map[string]models.Mission
	EngageEngine *engagement.Engine
}

func NewC2Server() *C2Server {
	return &C2Server{
		Hives:    make(map[string]models.Hive),
		Threats:  make(map[string]models.Threat),
		Missions: make(map[string]models.Mission),
	}
}

func (s *C2Server) AddThreat(t models.Threat) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Threats[t.ID] = t
}

func (s *C2Server) Engage(ctx context.Context, threatID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	threat, ok := s.Threats[threatID]
	if !ok {
		return "", fmt.Errorf("threat not found")
	}

	// Update engine with latest hives
	var hiveList []models.Hive
	for _, h := range s.Hives {
		hiveList = append(hiveList, h)
	}
	s.EngageEngine = &engagement.Engine{Hives: hiveList}

	mission, err := s.EngageEngine.PlanEngagement(threat)
	if err != nil {
		return "", err
	}

	s.Missions[mission.ID] = *mission
	fmt.Printf("Engagement initiated for threat %s, Mission: %s\n", threatID, mission.ID)

	return mission.ID, nil
}

func (s *C2Server) GetLiveStatus() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return map[string]interface{}{
		"hives":    s.Hives,
		"threats":  s.Threats,
		"missions": s.Missions,
	}
}

// GeoJSONExport provides data for Leaflet frontend
func (s *C2Server) GeoJSONExport() string {
	// Logic to generate GeoJSON for all active elements
	return "{\"type\": \"FeatureCollection\", \"features\": []}"
}
