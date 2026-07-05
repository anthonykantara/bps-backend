package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
	"ghost-hive/internal/hive"
	"time"
	"runtime"
)

func main() {
	fmt.Println("Starting GHOST HIVE ULTRA - Apocalypse Swarm Simulation (1 Million Threats)")

	server := c2.NewC2Server()
	server.GlobalStorage = 10000000 // 10 Million Global Reserve

	// 1. Initialize Large Scale Hive Network (100 Hives)
	for i := 0; i < 100; i++ {
		hiveID := fmt.Sprintf("HIVE-%d", i)
		h := models.Hive{
			ID:                hiveID,
			Location:          models.Coordinate{Lat: 48.0 + float64(i)*0.01, Lon: 2.0 + float64(i)*0.01},
			SafeCrashSite:     models.Coordinate{Lat: 47.0, Lon: 1.0, Alt: 0},
			Status:            models.HiveStatusActive,
			InterceptorsCount: 160,
			StorageCount:      10000,
			Interceptors:      make([]models.Interceptor, 160),
		}
		for j := 0; j < 160; j++ {
			h.Interceptors[j] = models.Interceptor{ID: fmt.Sprintf("I-%s-%d", hiveID, j)}
		}
		server.Hives[hiveID] = h
	}

	// 2. apocalypse Swarm (1,000,000 Threats)
	const numThreats = 1000000
	fmt.Printf("RADAR: Detecting APOCALYPSE SWARM (%d threats)...\n", numThreats)

	// Pre-allocate map to avoid excessive reallocations in simulation
	for i := 0; i < numThreats; i++ {
		if i % 100000 == 0 { fmt.Printf("Radar Ingest: %d/%d\n", i, numThreats) }
		t := models.Threat{
			ID:              fmt.Sprintf("T-%d", i),
			DroneType:       "SHARED-136",
			CurrentLocation: models.Coordinate{Lat: 49.0, Lon: 3.0, Alt: 1500},
			ProjectedTarget: models.Coordinate{Lat: 48.0, Lon: 2.0},
			Speed:           50.0,
			Altitude:        1500,
			Type:            models.ThreatSwarm,
		}
		server.AddThreat(t)
	}

	// 3. Mass Engagement with Auto-Reload Simulation
	fmt.Println("C2: Initiating Global Engagement...")

	start := time.Now()
	launched := 0
	for i := 0; i < numThreats; i++ {
		threatID := fmt.Sprintf("T-%d", i)
		_, err := server.Engage(context.Background(), threatID)
		if err == nil {
			launched++
		} else {
			// Trigger reload if empty
			for _, h := range server.Hives {
				if h.InterceptorsCount == 0 && h.StorageCount > 0 {
					ctrl := &hive.Controller{Hive: h}
					ctrl.Reload()
					server.Hives[h.ID] = ctrl.Hive
					break // Reload one and try next
				}
			}
		}

		if i % 100000 == 0 && i > 0 {
			fmt.Printf("Engagement Progress: %d missions launched...\n", launched)
		}
	}

	duration := time.Since(start)

	fmt.Printf("\n--- APOCALYPSE REPORT ---\n")
	fmt.Printf("Total Threats: %d\n", numThreats)
	fmt.Printf("Engagements Launched: %d\n", launched)
	fmt.Printf("Total Time: %v\n", duration)
	fmt.Printf("Interceptions per Second: %.2f\n", float64(launched)/duration.Seconds())

	stock := server.GetStockpileReport()
	fmt.Printf("Stockpile: Loaded=%d, Storage=%d, Total=%d\n", stock["loaded"], stock["storage"], stock["total"])

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Usage: %v MB\n", m.Alloc / 1024 / 1024)
}
