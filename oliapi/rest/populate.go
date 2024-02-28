package rest

import (
	"context"
	"oliapi/domain"
	"oliapi/ent"
)

func populateStaticData(ctx context.Context, ent *ent.Client) error {
	dbRoles, err := ent.Role.Query().All(ctx)
	roles := []domain.Role{
		domain.RoleAdmin,
		domain.RoleStaffAdmin,
	}

	if err != nil {
		return err //nolint:wrapcheck
	}

	for i := range roles {
		role := roles[i]
		found := false

		for j := range dbRoles {
			if dbRoles[j].Name == role.String() {
				found = true

				break
			}
		}

		if !found {
			_, err := ent.Role.Create().SetName(role.String()).Save(ctx)
			if err != nil {
				return err //nolint:wrapcheck
			}
		}
	}

	return nil
}
