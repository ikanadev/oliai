package rest

import (
	"context"
	"oliapi/domain"
)

func populateStaticData(ctx context.Context) error {
	roles := []domain.Role{
		domain.RoleAdmin,
		domain.RoleStaffAdmin,
	}

	for _ = range roles {
	}

	return nil
}
