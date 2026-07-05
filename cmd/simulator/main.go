package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
	"ghost-hive/internal/geo"
	"ghost-hive/internal/hive"
)

func main() {
	fmt.Println("--- GHOST HIVE EXTREME ENVIRONMENT SIMULATION ---")

	server := c2.NewC2Server()
	ctx := context.Background()

	// Setup baseline hives
	server.Hives["H-DESERT"] = models.Hive{ID: "H-DESERT", Status: models.HiveStatusActive, InterceptorsCount: 160, Location: models.Coordinate{Lat: 48, Lon: 2}, Environment: models.HiveEnvironment{Type: "DESERT", Temperature: 52}}
	server.Hives["H-ARCTIC"] = models.Hive{ID: "H-ARCTIC", Status: models.HiveStatusActive, InterceptorsCount: 160, Location: models.Coordinate{Lat: 48, Lon: 2}, Environment: models.HiveEnvironment{Type: "ARCTIC", Temperature: -35}}

	scenarios := []geo.WeatherCondition{
		{Type: geo.WeatherFog, Name: "HEAVY FOG", Visibility: 50, Temperature: 10, WindSpeed: 2},
		{Type: geo.WeatherSandstorm, Name: "SEVERE SANDSTORM", Visibility: 5, Temperature: 48, WindSpeed: 25},
		{Type: geo.WeatherSnow, Name: "BLIZZARD", Visibility: 200, Temperature: -15, WindSpeed: 30},
		{Type: geo.WeatherClear, Name: "EXTREME HEAT", Visibility: 10000, Temperature: 58, WindSpeed: 5},
	}

	for _, s := range scenarios {
		fmt.Printf("\n>>> SCENARIO: %s (Temp: %.1f C, Vis: %.1f m)\n", s.Name, s.Temperature, s.Visibility)
		geo.SetWeather(s)

		// Run Hive Env Management
		for _, h := range server.Hives {
			ctrl := &hive.Controller{Hive: h}
			ctrl.ManageEnvironment()
			server.Hives[h.ID] = ctrl.Hive
		}

		threat := models.Threat{ID: "T-" + s.Name, CurrentLocation: models.Coordinate{Lat: 48.5, Lon: 2.1}}
		server.AddThreat(threat)

		_, err := server.Engage(ctx, threat.ID)
		if err != nil {
			fmt.Printf("Engagement Result: BLOCKED (%v)\n", err)
		} else {
			fmt.Println("Engagement Result: AUTHORIZED")
		}
	}

	fmt.Println("\n--- EXTREME ENVIRONMENT SIMULATION COMPLETE ---")
}
