package request

import (
	"Hospital-Management-System/business/doctors"
)

type Doctors struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Fullname     string `json:"fullname"`
	Specialist   string `json:"specialist"`
	Address      string `json:"address"`
	DOB          string `json:"dob"`
	Phone_Number string `json:"phone_number"`
}
type DoctorLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Doctors) ToDomain() *doctors.Domain {
	return &doctors.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Fullname:     req.Fullname,
		Specialist:   req.Specialist,
		Address:      req.Address,
		DOB:          req.DOB,
		Phone_Number: req.Phone_Number,
	}
}
