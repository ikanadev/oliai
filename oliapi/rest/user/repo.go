package user

import (
	"context"
	"errors"
	"net/http"
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
func (u Repo) VerifyUser(ctx context.Context, email string, password string) (domain.User, error) {
	var respUser domain.User

	dbUser, err := u.ent.User.Query().Where(user.Email(email)).First(ctx)
	if err != nil {
		var entNotFound *ent.NotFoundError

		if errors.As(err, &entNotFound) {
			return respUser, utils.NewClientErr(http.StatusUnauthorized, utils.ErrEmailNotRegistered)
		}

		return respUser, utils.NewRestErr(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		return respUser, utils.NewClientErr(http.StatusUnauthorized, utils.ErrPasswordNotMatch)
	}

	respUser.ID = dbUser.ID
	respUser.FirstName = dbUser.FirstName
	respUser.LastName = dbUser.LastName
	respUser.Email = dbUser.Email

	return respUser, nil
}

// SaveUser implements repositories.UserRepository.
func (u Repo) SaveUser(ctx context.Context, data repository.SaveUserData) error {
	exists, err := u.ent.User.Query().Where(user.Email(data.Email)).Exist(ctx)
	if err != nil {
		return utils.NewRestErr(err)
	}

	if exists {
		return utils.NewClientErr(http.StatusConflict, utils.ErrEmailAlreadyRegistered)
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

	appUser = domain.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		TimeData: domain.TimeData{
			CreatedAt:  dbUser.CreatedAt,
			UpdatedAt:  dbUser.UpdatedAt,
			ArchivedAt: dbUser.ArchivedAt,
		},
	}

	return appUser, nil
}
