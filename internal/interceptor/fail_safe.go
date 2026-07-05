package interceptor

import (
	"fmt"
	"ghost-hive/internal/models"
	"sort"
)

type RiskType string

const (
	RiskHuman    RiskType = "HUMAN"
	RiskProperty RiskType = "PROPERTY"
	RiskAnimal   RiskType = "ANIMAL"
	RiskNone     RiskType = "OPEN_TERRAIN"
)

type CrashSite struct {
	Location models.Coordinate
	Risk     RiskType
	Priority int // Lower is safer/better (1: Open Terrain, 2: Animal, 3: Property, 4: Human)
}

type FailSafeHandler struct {
	Platform    *models.Interceptor
	Mission     models.Mission
	SafeZones   []CrashSite
}

// BingoBattery calculates the battery % needed to reach the safe crash site
func (f *FailSafeHandler) BingoBattery() float64 {
	// Simplified: 1% battery per 500m + 5% reserve
	return 10.0
}

func (f *FailSafeHandler) HandleThreatNotFound() {
	fmt.Printf("Interceptor %s: Threat not found at interception point. Initiating Search Mode.\n", f.Platform.ID)

	// 1. Expand Search Area: Higher altitude for better FOV
	fmt.Println("Increasing altitude to improve visual sensor sweep...")
	f.Platform.Location.Alt += 200

	// 2. Fly along route
	fmt.Printf("Scanning along projected target route: %+v\n", f.Mission.ProjectedTargetPt)
}

func (f *FailSafeHandler) ExecuteEmergencyLanding() {
	fmt.Printf("Interceptor %s: Battery critical or target search failed. Initiating safe landing.\n", f.Platform.ID)

	// Check if designated site is reachable
	if f.Platform.BatteryLevel >= f.BingoBattery() {
		fmt.Printf("Heading to designated Hive safe crash site at %+v\n", f.Mission.HiveSafeCrashSite)
	} else {
		fmt.Println("CRITICAL: Designated site unreachable. Scanning for immediate local alternatives...")

		bestSite := f.findBestCrashSite()
		if bestSite != nil {
			fmt.Printf("Heading to alternative crash site at %+v (Risk: %s)\n", bestSite.Location, bestSite.Risk)
		} else {
			fmt.Println("EMERGENCY: No safe zones found. Engaging AI-priority risk avoidance (1: Humans, 2: Property, 3: Animals).")
			fmt.Println("Executing final avoidance maneuver towards uninhabited coordinates.")
		}
	}

	fmt.Println("Executing controlled impact to minimize collateral damage.")
	f.Platform.Status = models.InterceptorStatusDestroyed
}

func (f *FailSafeHandler) findBestCrashSite() *CrashSite {
	if len(f.SafeZones) == 0 {
		return nil
	}

	// Priority: Open Terrain (RiskNone) is best
	sort.Slice(f.SafeZones, func(i, j int) bool {
		return f.SafeZones[i].Priority < f.SafeZones[j].Priority
	})

	return &f.SafeZones[0]
}
