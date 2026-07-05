package c2

import (
	"context"
	"fmt"
	"ghost-hive/internal/models"
	"time"
)

func (s *C2Server) StartAutoRedeploymentWorker(ctx context.Context) {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				var failedMissions []models.Mission
				s.mu.RLock()
				for _, mission := range s.Missions {
					if mission.Status == models.MissionFailed {
						failedMissions = append(failedMissions, mission)
					}
				}
				s.mu.RUnlock()

				for _, mission := range failedMissions {
					fmt.Printf("Orchestrator: Auto-Redeploying for threat %s (previous mission %s failed)\n", mission.ThreatID, mission.ID)
					_, err := s.Engage(ctx, mission.ThreatID)
					if err != nil {
						fmt.Printf("Orchestrator: Critical Failure - Unable to re-engage threat %s: %v\n", mission.ThreatID, err)
					} else {
						s.mu.Lock()
						m := s.Missions[mission.ID]
						m.Status = "RE-ENGAGED" // Mark old mission so we don't pick it up again
						s.Missions[mission.ID] = m
						s.mu.Unlock()
					}
				}
			}
		}
	}()
}
