package user

import (
	"context"
	"oliapi/domain"
	"oliapi/domain/repository"
	"oliapi/ent"
	"oliapi/ent/user"
	"oliapi/rest/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepo(ent *ent.Client) Repo {
	return Repo{ent: ent}
}

type Repo struct {
	ent *ent.Client
}

// VerifyUser implements repository.UserRepository.
func (Repo) VerifyUser(ctx context.Context, email string, password string) (domain.User, error) {
	panic("unimplemented")
}

// SaveUser implements repositories.UserRepository.
func (u Repo) SaveUser(ctx context.Context, data repository.SaveUserData) error {
	exists, err := u.ent.User.Query().Where(user.Email(data.Email)).Exist(ctx)
	if err != nil {
		return utils.NewRestErr(err)
	}

	if exists {
		return utils.NewClientErr(utils.HTTPConflict, "Correo ya registrado")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.NewRestErr(err)
	}

	_, err = u.ent.User.
		Create().
		SetEmail(data.Email).
		SetFirstName(data.FirstName).
		SetLastName(data.LastName).
		SetPassword(string(hashedPass)).
		Save(ctx)

	return utils.NewRestErr(err)
}

func (u Repo) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var appUser domain.User

	dbUser, err := u.ent.User.Query().WithRoles().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		return appUser, utils.NewRestErr(err)
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
