package hive

import (
	"fmt"
	"ghost-hive/internal/models"
)

type Controller struct {
	Hive models.Hive
}

func (c *Controller) SetStatus(status models.HiveStatus) {
	fmt.Printf("Hive %s switching to %s\n", c.Hive.ID, status)
	c.Hive.Status = status

	if status == models.HiveStatusStandby {
		fmt.Println("Powering up generator, charging interceptors, opening launch drawer...")
		// Simulate interceptor charging
		for i := range c.Hive.Interceptors {
			if c.Hive.Interceptors[i].BatteryLevel < 100 {
				c.Hive.Interceptors[i].Status = models.InterceptorStatusCharging
				c.Hive.Interceptors[i].BatteryLevel = 100 // Instant charge in standby for demo
			}
		}
	}
}

func (c *Controller) MonitorBMS() {
	// Battery Management System: Maintain health
	for i := range c.Hive.Interceptors {
		interceptor := &c.Hive.Interceptors[i]
		if interceptor.BatteryLevel < 20 {
			fmt.Printf("BMS: Low battery on interceptor %s, initiating trickle charge\n", interceptor.ID)
			interceptor.BatteryLevel += 1
		}
	}
}

func (c *Controller) LaunchInterceptor(id string) (*models.Interceptor, error) {
	for i := range c.Hive.Interceptors {
		if c.Hive.Interceptors[i].ID == id {
			c.Hive.Interceptors[i].Status = models.InterceptorStatusLaunched
			c.Hive.InterceptorsCount--
			return &c.Hive.Interceptors[i], nil
		}
	}
	return nil, fmt.Errorf("interceptor %s not found", id)
}
