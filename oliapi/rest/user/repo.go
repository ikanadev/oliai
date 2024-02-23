package user

import (
	"context"
	"oliapi/domain"
	"oliapi/ent"
	"oliapi/ent/user"

	"github.com/google/uuid"
)

type UserRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewUserRepo(ent *ent.Client) UserRepo {
	return UserRepo{ent: ent}
}

func (u UserRepo) GetUser(id uuid.UUID) (domain.User, error) {
	appUser := domain.User{}
	dbUser, err := u.ent.User.Query().WithRoles().Where(user.IDEQ(id)).Only(u.ctx)
	if err != nil {
		return appUser, err
	}
	roles := make([]domain.Role, len(dbUser.Edges.Roles))
	for i := range dbUser.Edges.Roles {
		roles[i] = domain.RoleFromSting(dbUser.Edges.Roles[i].Name)
	}
	appUser = domain.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Roles:     roles,
		TimeData: domain.TimeData{
			CreatedAt:  dbUser.CreatedAt,
			UpdatedAt:  dbUser.UpdatedAt,
			ArchivedAt: dbUser.ArchivedAt,
		},
	}
	return appUser, nil
}
