package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
)

func main() {
	fmt.Println("--- GHOST HIVE MIXED-PAYLOAD URBAN DEFENSE SIMULATION ---")

	server := c2.NewC2Server()
	ctx := context.Background()

	h := models.Hive{
		ID: "HIVE-URBAN-01",
		Status: models.HiveStatusActive,
		Location: models.Coordinate{Lat: 40.7128, Lon: -74.0060},
		Magazines: []models.Magazine{
			{ID: "MAG-K1", Type: models.TypeKinetic, InterceptorsCount: 40},
			{ID: "MAG-K2", Type: models.TypeKinetic, InterceptorsCount: 40},
			{ID: "MAG-K3", Type: models.TypeKinetic, InterceptorsCount: 40},
			{ID: "MAG-E1", Type: models.TypeExplosive, InterceptorsCount: 40},
		},
		InterceptorsCount: 160,
	}
	server.Hives[h.ID] = h

	fmt.Println("\n>>> SCENARIO A: Open Terrain / Low Risk")
	threatA := models.Threat{ID: "T-OPEN", CurrentLocation: models.Coordinate{Lat: 40.8, Lon: -74.1}, PayloadDetected: false}
	server.AddThreat(threatA)
	mID_A, _ := server.Engage(ctx, threatA.ID)
	fmt.Printf("C2: Mission %s engaged with type: %s\n", mID_A, server.Missions[mID_A].InterceptorType)

	fmt.Println("\n>>> SCENARIO B: Densely Populated Urban Zone / High Risk")
	threatB := models.Threat{
		ID: "T-URBAN-SHAHED",
		CurrentLocation: models.Coordinate{Lat: 40.72, Lon: -74.01},
		Altitude: 200,
		PayloadDetected: true,
	}
	server.AddThreat(threatB)
	mID_B, _ := server.Engage(ctx, threatB.ID)
	fmt.Printf("C2: Mission %s engaged with type: %s\n", mID_B, server.Missions[mID_B].InterceptorType)

	fmt.Println("\n>>> FINAL STOCKPILE REPORT")
	report := server.GetStockpileReport()
	fmt.Printf("Kinetic Loaded: %d\n", report["kinetic"].(map[string]int)["loaded"])
	fmt.Printf("Explosive Loaded: %d\n", report["explosive"].(map[string]int)["loaded"])

	fmt.Println("\n--- URBAN DEFENSE SIMULATION COMPLETE ---")
}
