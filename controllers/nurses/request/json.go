package request

import (
	"Hospital-Management-System/business/nurses"
)

type Nurses struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Fullname     string `json:"fullname"`
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	DOB          string `json:"dob"`
	Phone_Number string `json:"phone_number"`
}

type NurseLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Nurses) ToDomain() *nurses.Domain {
	return &nurses.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Fullname:     req.Fullname,
		Address:      req.Address,
		Gender:       req.Gender,
		DOB:          req.DOB,
		Phone_Number: req.Phone_Number,
	}
}
