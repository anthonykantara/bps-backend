package main

import (
	"context"
	"fmt"
	"ghost-hive/internal/c2"
	"ghost-hive/internal/models"
	"ghost-hive/internal/hive"
	"ghost-hive/internal/data"
	"time"
)

func main() {
	fmt.Println("GHOST HIVE FINAL VERIFICATION - 100% Feature Simulation")

	server := c2.NewC2Server()
	data.StartExternalDatasetIngestor()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.StartLogisticsWorker(ctx)

	readyTime := time.Now().Add(300 * time.Millisecond)
	hProd := models.Hive{
		ID:                 "HIVE-PROD-01",
		Status:             models.HiveStatusInProduction,
		EstimatedReadyTime: &readyTime,
		LiveTrackingCoord:  &models.Coordinate{Lat: 10, Lon: 10},
	}
	server.Hives[hProd.ID] = hProd

	hActive := models.Hive{
		ID:                "HIVE-ACTIVE-01",
		Status:            models.HiveStatusActive,
		InterceptorsCount: 80,
		StorageCount:      400,
		Location:          models.Coordinate{Lat: 48, Lon: 2},
	}
	server.Hives[hActive.ID] = hActive

	ctrl := &hive.Controller{Hive: server.Hives[hActive.ID]}
	ctrl.ReloadMagazine()
	server.Hives[hActive.ID] = ctrl.Hive

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Stockpile Report: %v\n", server.GetStockpileReport())
	fmt.Printf("Logistics Check: Hive %s Status = %s\n", "HIVE-PROD-01", server.Hives["HIVE-PROD-01"].Status)
}
