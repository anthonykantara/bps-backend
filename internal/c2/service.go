package c2

import (
	"context"
	"fmt"
	"ghost-hive/internal/engagement"
	"ghost-hive/internal/models"
	"sync"
)

type C2Server struct {
	mu            sync.RWMutex
	Hives         map[string]models.Hive
	Threats       map[string]models.Threat
	Missions      map[string]models.Mission
	GlobalStorage int
	EngageEngine  *engagement.Engine
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

	// Update Hive state (locally for simulation)
	hiveID := mission.HiveIDs[0]
	h := s.Hives[hiveID]
	h.InterceptorsCount--
	s.Hives[hiveID] = h

	return mission.ID, nil
}

func (s *C2Server) GetStockpileReport() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	loaded := 0
	storage := s.GlobalStorage
	for _, h := range s.Hives {
		loaded += h.InterceptorsCount
		storage += h.StorageCount
	}

	return map[string]int{
		"loaded":  loaded,
		"storage": storage,
		"total":   loaded + storage,
	}
}

func (s *C2Server) HandleMissionFailure(missionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	m, ok := s.Missions[missionID]
	if !ok {
		return
	}
	m.Status = models.MissionFailed
	s.Missions[missionID] = m

	fmt.Printf("C2: Mission %s failed. Triggering automated re-engagement for threat %s...\n", missionID, m.ThreatID)
	// New engagement will be triggered in the next loop or via a worker
}
