package rest

import (
	"context"
	"oliapi/domain"
	"oliapi/ent"
)

func populateStaticData(ent *ent.Client, ctx context.Context) error {
	roles := []domain.Role{
		domain.RoleAdmin,
		domain.RoleStaffAdmin,
	}
	dbRoles, err := ent.Role.Query().All(ctx)
	if err != nil {
		return err
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
				return err
			}
		}
	}
	return nil
}
