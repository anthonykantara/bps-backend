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
		for i := range c.Hive.Interceptors {
			c.Hive.Interceptors[i].Status = models.InterceptorStatusCharging
			c.Hive.Interceptors[i].BatteryLevel = 100
		}
	}
}

func (c *Controller) Reload() error {
	if c.Hive.StorageCount <= 0 {
		return fmt.Errorf("no storage remaining for hive %s", c.Hive.ID)
	}

	capacity := 160
	needed := capacity - c.Hive.InterceptorsCount
	if needed <= 0 {
		return nil
	}

	reloadAmount := needed
	if c.Hive.StorageCount < needed {
		reloadAmount = c.Hive.StorageCount
	}

	c.Hive.InterceptorsCount += reloadAmount
	c.Hive.StorageCount -= reloadAmount

	fmt.Printf("Hive %s: Reloaded %d interceptors from storage.\n", c.Hive.ID, reloadAmount)
	return nil
}

func (c *Controller) LaunchInterceptor(id string) (*models.Interceptor, error) {
	if c.Hive.InterceptorsCount <= 0 {
		return nil, fmt.Errorf("no interceptors loaded")
	}
	c.Hive.InterceptorsCount--
	return &models.Interceptor{Status: models.InterceptorStatusLaunched}, nil
}
