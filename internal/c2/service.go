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
	GlobalStorage map[models.InterceptorType]int
	EngageEngine  *engagement.Engine
	RBAC          *security.AccessControl
	IFF           *nato.IFFSystem
}

func NewC2Server() *C2Server {
	return &C2Server{
		Hives:    make(map[string]models.Hive),
		Threats:  make(map[string]models.Threat),
		Missions: make(map[string]models.Mission),
		GlobalStorage: map[models.InterceptorType]int{
			models.TypeKinetic: 10000,
			models.TypeExplosive: 5000,
		},
		RBAC:     &security.AccessControl{ActiveUser: models.User{ID: "CMD-01", Role: models.RoleCommander}},
		IFF:      &nato.IFFSystem{FriendlyCodes: map[string]bool{"NATO-X-RAY": true}},
	}
}

func (s *C2Server) Engage(ctx context.Context, threatID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := s.RBAC.CanEngage("SINGLE_ENGAGE"); err != nil { return "", err }

	threat, ok := s.Threats[threatID]
	if !ok { return "", fmt.Errorf("threat not found") }

	isHostile, reason := s.IFF.FilterThreat(threat)
	if !isHostile { return "", fmt.Errorf("IFF BLOCK: %s", reason) }

	var hiveList []models.Hive
	for _, h := range s.Hives { hiveList = append(hiveList, h) }
	s.EngageEngine = &engagement.Engine{Hives: hiveList}

	mission, err := s.EngageEngine.PlanEngagement(threat)
	if err != nil { return "", err }

	// Launch from hive
	hiveID := mission.HiveIDs[0]
	h := s.Hives[hiveID]

	// Update Hive local state
	for i := range h.Magazines {
		if h.Magazines[i].Type == mission.InterceptorType && h.Magazines[i].InterceptorsCount > 0 {
			h.Magazines[i].InterceptorsCount--
			h.InterceptorsCount--
			break
		}
	}
	s.Hives[hiveID] = h

	s.Missions[mission.ID] = *mission
	return mission.ID, nil
}

func (s *C2Server) GetStockpileReport() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	loadedK, loadedE := 0, 0
	storageK, storageE := s.GlobalStorage[models.TypeKinetic], s.GlobalStorage[models.TypeExplosive]

	for _, h := range s.Hives {
		for _, m := range h.Magazines {
			if m.Type == models.TypeKinetic { loadedK += m.InterceptorsCount }
			if m.Type == models.TypeExplosive { loadedE += m.InterceptorsCount }
		}
		storageK += h.StorageCount[models.TypeKinetic]
		storageE += h.StorageCount[models.TypeExplosive]
	}

	return map[string]interface{}{
		"kinetic": map[string]int{"loaded": loadedK, "storage": storageK},
		"explosive": map[string]int{"loaded": loadedE, "storage": storageE},
		"total_ready": loadedK + loadedE,
	}
}

func (s *C2Server) AddThreat(t models.Threat) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Threats[t.ID] = t
}

func (s *C2Server) HandleMissionFailure(missionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if m, ok := s.Missions[missionID]; ok {
		m.Status = models.MissionFailed
		s.Missions[missionID] = m
	}
}
