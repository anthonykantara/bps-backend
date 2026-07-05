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
			fmt.Printf("Hive %s: Extreme cold (%.1f C). Increasing heater load to prevent magazine freeze.\n", c.Hive.ID, env.Temperature)
			env.HeaterLoad = 1.0
			env.CoolingLoad = 0.0
		}
	case "DESERT":
		if env.Temperature > 45 {
			fmt.Printf("Hive %s: Extreme heat (%.1f C). Activating active liquid cooling for magazines.\n", c.Hive.ID, env.Temperature)
			env.CoolingLoad = 1.0
			env.HeaterLoad = 0.0
		}
	}
}

func (c *Controller) GetEnvTelemetry() models.HiveEnvironment {
	return c.Hive.Environment
}
