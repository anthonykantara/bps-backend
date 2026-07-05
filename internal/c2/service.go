package c2

import (
	"context"
	"fmt"
	"ghost-hive/internal/engagement"
	"ghost-hive/internal/models"
	"ghost-hive/internal/security"
	"ghost-hive/internal/nato"
	"sync"
)

type C2Server struct {
	mu            sync.RWMutex
	Hives         map[string]models.Hive
	Threats       map[string]models.Threat
	Missions      map[string]models.Mission
	GlobalStorage int
	EngageEngine  *engagement.Engine
	RBAC          *security.AccessControl
	IFF           *nato.IFFSystem
}

func NewC2Server() *C2Server {
	return &C2Server{
		Hives:    make(map[string]models.Hive),
		Threats:  make(map[string]models.Threat),
		Missions: make(map[string]models.Mission),
		RBAC:     &security.AccessControl{ActiveUser: models.User{ID: "CMD-01", Role: models.RoleCommander}},
		IFF:      &nato.IFFSystem{FriendlyCodes: map[string]bool{"NATO-X-RAY": true}},
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

	// 1. RBAC Check
	if _, err := s.RBAC.CanEngage("SINGLE_ENGAGE"); err != nil {
		return "", err
	}

	threat, ok := s.Threats[threatID]
	if !ok {
		return "", fmt.Errorf("threat not found")
	}

	// 2. IFF Check
	isHostile, reason := s.IFF.FilterThreat(threat)
	if !isHostile {
		return "", fmt.Errorf("IFF BLOCK: target %s identified as %s", threatID, reason)
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

	mission.AuthorizedBy = s.RBAC.ActiveUser.ID
	s.Missions[mission.ID] = *mission

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
	return map[string]int{"loaded": loaded, "storage": storage, "total": loaded + storage}
}

func (s *C2Server) HandleMissionFailure(missionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if m, ok := s.Missions[missionID]; ok {
		m.Status = models.MissionFailed
		s.Missions[missionID] = m
	}
}
