package user

import (
	"database/sql"
	"errors"
	"net/http"
	"oliapi/db"
	"oliapi/domain"
	"oliapi/domain/repository"
	"oliapi/rest/utils"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

type Repo struct {
	db *sqlx.DB
}

// VerifyUser implements repository.UserRepository.
func (u Repo) VerifyUser(email string, password string) (domain.User, error) {
	var (
		respUser domain.User
		dbUser   db.User
	)

	err := u.db.Get(&dbUser, "select * from users where email = $1;", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return respUser, echo.NewHTTPError(http.StatusNotFound, utils.ErrEmailNotRegistered)
		}

		return respUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return respUser, echo.NewHTTPError(http.StatusUnauthorized, utils.ErrPasswordIncorrect)
		}

		return respUser, err
	}

	respUser.ID = dbUser.ID
	respUser.Email = dbUser.Email
	respUser.FirstName = dbUser.FirstName
	respUser.LastName = dbUser.LastName

	return respUser, nil
}

// SaveUser implements repositories.UserRepository.
func (u Repo) SaveUser(data repository.SaveUserData) error {
	var mails []string
	now := time.Now()

	err := u.db.Select(&mails, "select email from users where email = $1;", data.Email)
	if err != nil {
		return err
	}

	if len(mails) > 0 {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrEmailAlreadyRegistered)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = u.db.Exec(`
insert into users
(id, email, first_name, last_name, password, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6, $7);`,
		uuid.New(), data.Email, data.FirstName, data.LastName, hashedPassword, now, now,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u Repo) GetUser(id uuid.UUID) (domain.User, error) {
	var appUser domain.User

	err := u.db.Get(&appUser, "select * from users where id = $1;", id)
	if err != nil {
		return appUser, err
	}

	return appUser, nil
}
