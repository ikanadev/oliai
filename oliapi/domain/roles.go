package domain

type Role int

const (
	RoleAdmin Role = iota
	StaffAdmin
)

func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case StaffAdmin:
		return "admin_staff"
	}
	return "unknown"
}
