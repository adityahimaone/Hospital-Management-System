package nurses

import "time"

type Domain struct {
	ID           int
	Username     string
	Password     string
	Fullname     string
	Address      string
	Gender      string
	DOB          string
	Phone_Number string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	AllNurse() ([]Domain, error)
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	NurseByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllNurse() ([]Domain, error)
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	NurseByID(id int) (Domain, error)
	Delete(id int) (string, error)
}