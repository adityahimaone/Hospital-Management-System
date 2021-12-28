package doctors

import "time"

type Domain struct {
	ID           int
	Username     string
	Password     string
	Name         string
	Fullname     string
	Specialist   string
	Address      string
	Phone_Number string
	DOB          string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	AllDoctor() ([]Domain, error)
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	DoctorByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllDoctor() ([]Domain, error)
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	DoctorByID(id int) (Domain, error)
	Delete(id int) (string, error)
}
