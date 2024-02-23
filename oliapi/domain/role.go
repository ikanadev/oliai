package domain

type Role int

const (
	RoleAdmin Role = iota
	RoleStaffAdmin
)

func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleStaffAdmin:
		return "admin_staff"
	}
	return "unknown"
}

func RoleFromSting(s string) Role {
	switch s {
	case "admin":
		return RoleAdmin
	case "admin_staff":
		return RoleStaffAdmin
	}
	return RoleStaffAdmin
}
