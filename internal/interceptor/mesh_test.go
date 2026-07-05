package interceptor

import (
	"ghost-hive/internal/models"
	"testing"
)

func TestCoordinateTarget(t *testing.T) {
	node := &MeshNode{
		ID:    "INT-1",
		Peers: make(map[string]*MeshNode),
	}

	swarm := []models.Threat{
		{ID: "T1"},
		{ID: "T2"},
	}

	// Peer 1 has already claimed T1
	node.Peers["INT-2"] = &MeshNode{ID: "INT-2", CurrentTarget: "T1"}

	target := node.CoordinateTarget(swarm)

	if target != "T2" {
		t.Errorf("Expected target T2, got %s", target)
	}
}
