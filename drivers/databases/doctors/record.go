package doctors

import (
	"Hospital-Management-System/business/doctors"
	"time"

	"gorm.io/gorm"
)

type Doctors struct {
	gorm.Model
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	Password     string
	Fullname     string
	Specialist   string
	Address      string
	Phone_Number string
	DOB          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func toDomain(org Doctors) doctors.Domain {
	return doctors.Domain{
		ID:           org.ID,
		Username:     org.Username,
		Password:     org.Password,
		Fullname:     org.Fullname,
		Phone_Number: org.Phone_Number,
		Specialist:   org.Specialist,
		Address:      org.Address,
		DOB:          org.DOB,
		CreatedAt:    org.CreatedAt,
		UpdatedAt:    org.UpdatedAt,
	}
}
func toDomainUpdate(upd Doctors) doctors.Domain {
	return doctors.Domain{
		ID:           upd.ID,
		Username:     upd.Username,
		Password:     upd.Password,
		Fullname:     upd.Fullname,
		Phone_Number: upd.Phone_Number,
		Specialist:   upd.Specialist,
		Address:      upd.Address,
		DOB:          upd.DOB,
		CreatedAt:    upd.CreatedAt,
		UpdatedAt:    upd.UpdatedAt,
	}
}
func fromDomain(domain doctors.Domain) Doctors {
	return Doctors{
		ID:           domain.ID,
		Username:     domain.Username,
		Password:     domain.Password,
		Fullname:     domain.Fullname,
		Phone_Number: domain.Phone_Number,
		Specialist:   domain.Specialist,
		Address:      domain.Address,
		DOB:          domain.DOB,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
func toDomainList(data []Doctors) []doctors.Domain {
	result := []doctors.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
