package domain

type Role string

const (
	RoleAdmin Role = "admin"
	//	RoleStaff users are OLIAI people that is in charge of manage company's bots.
	//	Admin will assign companies to them.
	RoleStaff Role = "staff"
	//	RoleUser users are the company's designed "admin" from their bots.
	//	So they can UPDATE company's bots, categories, documents.
	RoleUser Role = "user"
)

func RoleFromSting(s string) Role {
	switch s {
	case "admin":
		return RoleAdmin
	case "staff":
		return RoleStaff
	case "user":
		return RoleUser
	}

	return RoleUser
}
