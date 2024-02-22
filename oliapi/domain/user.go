package domain

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Roles    []Role `json:"roles"`
	TimeData
}
