package admins

import "time"

type Domain struct {
	ID        int
	Username  string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Login(username, password string) (Domain, error)
}

type Repository interface {
	Login(username, password string) (Domain, error)
}
