package engagement

import (
	"fmt"
	"ghost-hive/internal/models"
)

// WaveAnalysis identifies swarm intent (e.g., distraction vs. main strike)
func AnalyzeWave(swarm []models.Threat) string {
	if len(swarm) > 100 && swarm[0].Speed < 40 {
		return "PROBABLE DISTRACTION WAVE - Conserve high-performance interceptors."
	}
	return "CRITICAL STRIKE - Maximize response."
}

type SupplyChainManager struct {
	StockpileThreshold int
}

func (s *SupplyChainManager) CheckStockpile(hive models.Hive) {
	total := hive.InterceptorsCount + hive.StorageCount
	if total < s.StockpileThreshold {
		fmt.Printf("SUPPLY CHAIN: Hive %s low on stock (%d). Automated transport mission scheduled.\n", hive.ID, total)
	}
}
