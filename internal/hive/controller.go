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
}

func (c *Controller) ReloadMagazine(t models.InterceptorType) error {
	const magazineSize = 40
	if c.Hive.StorageCount[t] < magazineSize {
		return fmt.Errorf("insufficient %s storage for magazine reload in hive %s", t, c.Hive.ID)
	}

	if len(c.Hive.Magazines) >= 4 {
		return fmt.Errorf("hive %s magazine slots full (4 max)", c.Hive.ID)
	}

	c.Hive.InterceptorsCount += magazineSize
	c.Hive.StorageCount[t] -= magazineSize

	mag := models.Magazine{
		ID: fmt.Sprintf("MAG-%s-%d", c.Hive.ID, len(c.Hive.Magazines)),
		Type: t,
		InterceptorsCount: magazineSize,
	}
	c.Hive.Magazines = append(c.Hive.Magazines, mag)

	fmt.Printf("Hive %s: Loaded %s magazine. Total Loaded: %d\n", c.Hive.ID, t, c.Hive.InterceptorsCount)
	return nil
}

func (c *Controller) LaunchInterceptor(t models.InterceptorType) (*models.Interceptor, error) {
	// Find first available magazine of requested type
	for i := range c.Hive.Magazines {
		if c.Hive.Magazines[i].Type == t && c.Hive.Magazines[i].InterceptorsCount > 0 {
			c.Hive.Magazines[i].InterceptorsCount--
			c.Hive.InterceptorsCount--

			// Eject if empty
			if c.Hive.Magazines[i].InterceptorsCount == 0 {
				c.Hive.Magazines = append(c.Hive.Magazines[:i], c.Hive.Magazines[i+1:]...)
			}

			return &models.Interceptor{Type: t, Status: models.InterceptorStatusLaunched}, nil
		}
	}

	return nil, fmt.Errorf("no %s interceptors available in hive %s", t, c.Hive.ID)
}

func (c *Controller) ManageEnvironment() {
	env := &c.Hive.Environment
	switch env.Type {
	case "AMPHIBIOUS":
		if env.Pressure > 1.5 { env.SealIntegrity = 0.99 }
	case "ARCTIC":
		if env.Temperature < -20 { env.HeaterLoad = 1.0 }
	case "DESERT":
		if env.Temperature > 45 { env.CoolingLoad = 1.0 }
	}
}
