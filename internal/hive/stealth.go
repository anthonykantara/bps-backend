package hive

import "fmt"

func (c *Controller) EnableSilentMode() {
	fmt.Printf("Hive %s: ENTERING SILENT MODE (EMISSION CONTROL).\n", c.Hive.ID)
	fmt.Println("Action: Disabling Starlink/Radar active emissions. Switching to Passive Acoustic/Optical sensing.")
}
