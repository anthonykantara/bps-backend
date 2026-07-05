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

func (c *Controller) ReloadMagazine() error {
	const magazineSize = 40
	if c.Hive.StorageCount < magazineSize {
		return fmt.Errorf("insufficient storage for full magazine reload in hive %s", c.Hive.ID)
	}

	if c.Hive.InterceptorsCount + magazineSize > 160 {
		return fmt.Errorf("hive %s exceeds interceptor capacity (160 max)", c.Hive.ID)
	}

	c.Hive.InterceptorsCount += magazineSize
	c.Hive.StorageCount -= magazineSize

	mag := models.Magazine{
		ID: fmt.Sprintf("MAG-%s-%d", c.Hive.ID, len(c.Hive.Magazines)),
		InterceptorsCount: magazineSize,
	}
	c.Hive.Magazines = append(c.Hive.Magazines, mag)

	fmt.Printf("Hive %s: Magazine-based reload complete (40 units). Total Loaded: %d\n", c.Hive.ID, c.Hive.InterceptorsCount)
	return nil
}

func (c *Controller) LaunchInterceptor(id string) (*models.Interceptor, error) {
	if c.Hive.InterceptorsCount <= 0 {
		return nil, fmt.Errorf("no interceptors loaded")
	}
	c.Hive.InterceptorsCount--

	// Deplete current magazine
	if len(c.Hive.Magazines) > 0 {
		idx := len(c.Hive.Magazines)-1
		c.Hive.Magazines[idx].InterceptorsCount--
		if c.Hive.Magazines[idx].InterceptorsCount == 0 {
			c.Hive.Magazines = c.Hive.Magazines[:idx] // Eject empty magazine
		}
	}

	return &models.Interceptor{Status: models.InterceptorStatusLaunched}, nil
}
