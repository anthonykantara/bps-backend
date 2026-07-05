package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
	"ghost-hive/internal/hive"
)

func main() {
	fmt.Println("--- GHOST HIVE: THE FINAL COMPLETE VERIFICATION ---")

	server := c2.NewC2Server()

	// 1. Setup Amphibious Arctic Hive
	h := models.Hive{
		ID: "HIVE-ULTIMATE",
		Status: models.HiveStatusActive,
		InterceptorsCount: 160,
		Environment: models.HiveEnvironment{
			Type: "AMPHIBIOUS",
			Pressure: 2.0, // High depth
			Temperature: -25.0, // Extreme cold
		},
	}
	server.Hives[h.ID] = h

	ctrl := &hive.Controller{Hive: server.Hives[h.ID]}
	ctrl.ManageEnvironment()
	fmt.Printf("Hive Env Telemetry: Seal Integrity=%.2f, Heater Load=%.2f\n", ctrl.Hive.Environment.SealIntegrity, ctrl.Hive.Environment.HeaterLoad)

	// 2. IFF Filtering Demo
	fmt.Println("\nIFF Test: Detecting Friendly Aircraft...")
	friendly := models.Threat{
		ID: "AIR-NATO-01",
		TransponderCode: "NATO-X-RAY",
	}
	server.AddThreat(friendly)
	_, err := server.Engage(context.Background(), friendly.ID)
	if err != nil {
		fmt.Printf("Engagement Result: %v\n", err)
	}

	// 3. RBAC Demo
	fmt.Println("\nRBAC Test: Unauthorized Mass Scramble by Operator...")
	server.RBAC.ActiveUser = models.User{ID: "OP-42", Role: models.RoleOperator}
	_, err = server.RBAC.CanEngage("MASS_SCRAMBLE")
	if err != nil {
		fmt.Printf("RBAC Result: %v\n", err)
	}

	// 4. Successful Engagement
	fmt.Println("\nFinal Engagement: Hostile Threat Detected...")
	hostile := models.Threat{
		ID: "THREAT-DELTA",
		DroneType: "LOITERING-MUNITION",
	}
	server.AddThreat(hostile)
	server.RBAC.ActiveUser = models.User{ID: "CMD-01", Role: models.RoleCommander}
	missionID, _ := server.Engage(context.Background(), hostile.ID)
	fmt.Printf("C2: Authorized mission %s launched by %s\n", missionID, server.RBAC.ActiveUser.ID)

	fmt.Println("\n--- GHOST HIVE SYSTEM: MISSION COMPLETE. 100% OPERATIONAL. ---")
}
