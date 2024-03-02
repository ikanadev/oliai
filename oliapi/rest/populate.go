package rest

import (
	"oliapi/domain"

	"github.com/google/uuid"
)

func (s Server) Populate() {
	roles := []domain.Role{
		domain.RoleAdmin,
		domain.RoleStaff,
		domain.RoleUser,
	}
	sqlStatement := `
insert into roles (id, name)
select $1, $2
where not exists (
	select 1 from roles where name = $3
);
`

	for i := range roles {
		_, err := s.db.Exec(sqlStatement, uuid.New(), roles[i], roles[i])
		panicIfError(err)
	}
}
