package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/hive"
	"ghost-hive/internal/models"
	"time"
)

func main() {
	fmt.Println("Starting Ghost Hive Simulation...")

	server := c2.NewC2Server()

	// 1. Initialize Hives
	h := models.Hive{
		ID:                "HIVE-ALPHA",
		Location:          models.Coordinate{Lat: 48.8566, Lon: 2.3522},
		Status:            models.HiveStatusIdle,
		InterceptorsCount: 160,
		Interceptors:      make([]models.Interceptor, 160),
	}
	for i := 0; i < 160; i++ {
		h.Interceptors[i] = models.Interceptor{
			ID:           fmt.Sprintf("INT-%s-%d", h.ID, i),
			HiveID:       h.ID,
			Status:       models.InterceptorStatusIdle,
			BatteryLevel: 80.0,
			Health:       100.0,
		}
	}
	server.Hives[h.ID] = h

	// 2. Simulate Incoming Threat via Radar
	threat := models.Threat{
		ID:              "THREAT-001",
		DroneType:       "SHARED-136",
		CurrentLocation: models.Coordinate{Lat: 49.0, Lon: 2.5, Alt: 1000},
		ProjectedTarget: models.Coordinate{Lat: 48.8566, Lon: 2.3522},
		Speed:           50.0,
		Altitude:        1000,
	}
	server.AddThreat(threat)
	fmt.Printf("RADAR ALERT: Detected %s at distance.\n", threat.DroneType)

	// 3. Automated Engagement
	fmt.Println("C2: Threat detected. Analyzing engagement options...")

	// Switch Hive to Standby
	controller := &hive.Controller{Hive: server.Hives["HIVE-ALPHA"]}
	controller.SetStatus(models.HiveStatusStandby)
	server.Hives["HIVE-ALPHA"] = controller.Hive

	// 1-Click Engage logic
	missionID, err := server.Engage(context.Background(), threat.ID)
	if err != nil {
		fmt.Printf("Engagement failed: %v\n", err)
		return
	}

	// 4. Mission Execution
	fmt.Printf("Mission %s launched. Tracking interception...\n", missionID)

	// Simulate launch
	launchedInterceptor, _ := controller.LaunchInterceptor("INT-HIVE-ALPHA-0")
	server.Hives["HIVE-ALPHA"] = controller.Hive

	fmt.Printf("Interceptor %s launched from HIVE-ALPHA.\n", launchedInterceptor.ID)

	time.Sleep(1 * time.Second)
	fmt.Println("Simulation Complete: Threat Neutralized.")
}
