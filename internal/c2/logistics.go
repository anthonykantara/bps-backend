package c2

import (
	"context"
	"fmt"
	"ghost-hive/internal/models"
	"time"
)

func (s *C2Server) StartLogisticsWorker(ctx context.Context) {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.mu.Lock()
				for id, hive := range s.Hives {
					if hive.Status == models.HiveStatusInProduction && hive.EstimatedReadyTime != nil {
						if time.Now().After(*hive.EstimatedReadyTime) {
							fmt.Printf("Logistics: Hive %s production complete. Switching to Transport.\n", id)
							hive.Status = models.HiveStatusInTransport
							s.Hives[id] = hive
						}
					}

					if hive.Status == models.HiveStatusInTransport && hive.LiveTrackingCoord != nil {
						// Simulate transport movement towards target destination
						hive.LiveTrackingCoord.Lat += 0.001
						s.Hives[id] = hive
					}
				}
				s.mu.Unlock()
			}
		}
	}()
}
