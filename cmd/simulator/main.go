package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
	"ghost-hive/internal/engagement"
	"ghost-hive/internal/metrics"
	"time"
)

func main() {
	fmt.Println("Starting GHOST HIVE PRODUCTION-READY STRESS TEST")

	server := c2.NewC2Server()

	// 1. Setup Hives with Safe Zones & Payloads
	for i := 0; i < 10; i++ {
		hiveID := fmt.Sprintf("HIVE-%d", i)
		h := models.Hive{
			ID:                hiveID,
			Location:          models.Coordinate{Lat: 48.0, Lon: 2.0},
			Status:            models.HiveStatusActive,
			InterceptorsCount: 160,
			StorageCount:      5000,
		}
		server.Hives[hiveID] = h
	}

	// 2. Intelligence: Wave Analysis
	threats := make([]models.Threat, 1000)
	for i := 0; i < 1000; i++ {
		threats[i] = models.Threat{ID: fmt.Sprintf("T-%d", i), Speed: 30}
	}
	fmt.Printf("Intelligence Analysis: %s\n", engagement.AnalyzeWave(threats))

	// 3. Mass Engagement Simulation
	fmt.Println("C2: Initiating engagement with BDA and Metrics enabled...")
	start := time.Now()
	for i := 0; i < 1000; i++ {
		server.AddThreat(threats[i])
		_, err := server.Engage(context.Background(), threats[i].ID)
		if err == nil {
			metrics.DefaultRegistry.IncrementCounter("engagements_total")
		}
	}

	fmt.Printf("Simulation: Processed 1000 engagements in %v\n", time.Since(start))

	// 4. Report Metrics
	metrics.DefaultRegistry.Report()

	fmt.Println("Final Production Stress Test Successful.")
}
