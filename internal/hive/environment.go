package hive

import (
	"fmt"
	"ghost-hive/internal/models"
)

func (c *Controller) ManageEnvironment() {
	env := &c.Hive.Environment

	switch env.Type {
	case "AMPHIBIOUS":
		if env.Pressure > 1.5 {
			fmt.Printf("Hive %s: High external pressure detected (%.2f bar). Reinforcing seals.\n", c.Hive.ID, env.Pressure)
			env.SealIntegrity = 0.99
		}
	case "ARCTIC":
		if env.Temperature < -20 {
			fmt.Printf("Hive %s: Extreme cold detected (%.1f C). Increasing heater load to prevent magazine freeze.\n", c.Hive.ID, env.Temperature)
			env.HeaterLoad = 0.8
		}
	}
}

func (c *Controller) GetEnvTelemetry() models.HiveEnvironment {
	return c.Hive.Environment
}
