package security

import (
	"fmt"
	"ghost-hive/internal/models"
)

type AccessControl struct {
	ActiveUser models.User
}

func (a *AccessControl) CanEngage(action string) (bool, error) {
	switch a.ActiveUser.Role {
	case models.RoleCommander, models.RoleAdmin:
		return true, nil
	case models.RoleOperator:
		if action == "MASS_SCRAMBLE" {
			return false, fmt.Errorf("insufficient permissions: Commander role required for mass scramble")
		}
		return true, nil
	default:
		return false, fmt.Errorf("unauthorized")
	}
}

func (a *AccessControl) LogAction(action string) {
	fmt.Printf("RBAC: User %s (%s) authorized for %s\n", a.ActiveUser.ID, a.ActiveUser.Role, action)
}
