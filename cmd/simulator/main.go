package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
	"ghost-hive/internal/geo"
	"time"
)

func main() {
	fmt.Println("--- GHOST HIVE MULTI-VARIABLE SIMULATION SUITE ---")

	server := c2.NewC2Server()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the auto-redeployment worker
	server.StartAutoRedeploymentWorker(ctx)

	// Setup baseline hives
	for i := 0; i < 5; i++ {
		server.Hives[fmt.Sprintf("H-%d", i)] = models.Hive{
			ID: fmt.Sprintf("H-%d", i),
			Status: models.HiveStatusActive,
			InterceptorsCount: 160,
			Location: models.Coordinate{Lat: 48, Lon: 2},
		}
	}

	// 1. Weather Scenarios
	weatherScenarios := []geo.WeatherCondition{
		{Name: "CLEAR SKIES", WindSpeed: 5, Visibility: 10000, IsJamming: false},
		{Name: "HIGH WIND (GALE)", WindSpeed: 45, Visibility: 5000, IsJamming: false},
		{Name: "HURRICANE", WindSpeed: 120, Visibility: 50, IsJamming: false},
	}

	for _, s := range weatherScenarios {
		fmt.Printf("\n>>> WEATHER TEST: %s\n", s.Name)
		geo.SetWeather(s)

		threat := models.Threat{ID: "T-" + s.Name, CurrentLocation: models.Coordinate{Lat: 48.5, Lon: 2.1}, DroneType: "UAV"}
		server.AddThreat(threat)
		_, err := server.Engage(ctx, threat.ID)
		if err != nil {
			fmt.Printf("Result: %v\n", err)
		} else {
			fmt.Println("Result: Success (Authorized)")
		}
	}

	// 2. Failure & Auto-Redeployment Test
	fmt.Println("\n>>> FAILURE TEST: Manual Interceptor/Mission Crash")
	geo.SetWeather(geo.WeatherCondition{Name: "CLEAR", WindSpeed: 5, Visibility: 10000})

	threatID := "THREAT-CRASH-TEST"
	server.AddThreat(models.Threat{ID: threatID, CurrentLocation: models.Coordinate{Lat: 48.5, Lon: 2.1}, DroneType: "UAV"})

	missionID, _ := server.Engage(ctx, threatID)
	fmt.Printf("Initial Mission %s launched.\n", missionID)

	// Simulate mission failure (crash)
	fmt.Println("CRITICAL: Mission failure detected (interceptor crashed). Triggering HandleMissionFailure...")
	server.HandleMissionFailure(missionID)

	// Wait for worker to re-engage
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n--- SIMULATION SUITE COMPLETE ---")
}
