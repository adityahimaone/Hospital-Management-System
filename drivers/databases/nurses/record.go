package nurses

import (
	"Hospital-Management-System/business/nurses"
	"time"

	"gorm.io/gorm"
)

type Nurses struct {
	gorm.Model
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	Password     string
	Fullname     string
	Address      string
	Gender       string
	DOB          string
	Phone_Number string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func toDomain(org Nurses) nurses.Domain {
	return nurses.Domain{
		ID:           org.ID,
		Username:     org.Username,
		Password:     org.Password,
		Fullname:     org.Fullname,
		Address:      org.Address,
		Gender:       org.Gender,
		DOB:          org.DOB,
		Phone_Number: org.Phone_Number,
		CreatedAt:    org.CreatedAt,
		UpdatedAt:    org.UpdatedAt,
	}
}

func toDomainUpdate(upd Nurses) nurses.Domain {
	return nurses.Domain{
		ID:           upd.ID,
		Username:     upd.Username,
		Password:     upd.Password,
		Fullname:     upd.Fullname,
		Address:      upd.Address,
		Gender:       upd.Gender,
		DOB:          upd.DOB,
		Phone_Number: upd.Phone_Number,
		CreatedAt:    upd.CreatedAt,
		UpdatedAt:    upd.UpdatedAt,
	}
}
func fromDomain(domain nurses.Domain) Nurses {
	return Nurses{
		ID:           domain.ID,
		Username:     domain.Username,
		Password:     domain.Password,
		Fullname:     domain.Fullname,
		Address:      domain.Address,
		Gender:       domain.Gender,
		DOB:          domain.DOB,
		Phone_Number: domain.Phone_Number,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
func toDomainList(data []Nurses) []nurses.Domain {
	result := []nurses.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
