package engagement

import (
	"ghost-hive/internal/models"
	"testing"
)

func TestPlanEngagement(t *testing.T) {
	hives := []models.Hive{
		{
			ID:                "HIVE-1",
			Location:          models.Coordinate{Lat: 10, Lon: 10},
			Status:            models.HiveStatusActive,
			InterceptorsCount: 10,
			Magazines: []models.Magazine{
				{ID: "M1", Type: models.TypeKinetic, InterceptorsCount: 10},
			},
		},
	}
	engine := &Engine{Hives: hives}
	threat := models.Threat{
		ID:              "T1",
		CurrentLocation: models.Coordinate{Lat: 11, Lon: 11},
	}

	mission, err := engine.PlanEngagement(threat)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if mission.HiveIDs[0] != "HIVE-1" {
		t.Errorf("Expected HIVE-1, got %s", mission.HiveIDs[0])
	}
}
