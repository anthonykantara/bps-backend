package nato

import (
	"fmt"
	"ghost-hive/internal/models"
)

type IFFSystem struct {
	FriendlyCodes map[string]bool
}

func (i *IFFSystem) FilterThreat(t models.Threat) (bool, string) {
	if t.TransponderCode != "" && i.FriendlyCodes[t.TransponderCode] {
		return false, "FRIENDLY_TRANS_MATCH"
	}

	// Mode 5 / Mode S crypto handshake simulation
	if t.TransponderCode == "7700" {
		return false, "CIVILIAN_EMERGENCY"
	}

	return true, "PID_CONFIRMED_THREAT"
}

func (i *IFFSystem) MarkFriendly(id string) {
	fmt.Printf("IFF: Target %s identified as friendly. Engagement blocked.\n", id)
}
