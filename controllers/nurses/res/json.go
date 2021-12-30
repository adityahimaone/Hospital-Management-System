package response

import (
	"Hospital-Management-System/business/nurses"
	"time"
)

type NurseRegisterRespons struct {
	Message      string    `json:"message"`
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Fullname     string    `json:"fullname"`
	Address      string    `json:"address"`
	Gender       string    `json:"gender"`
	DOB          string    `json:"dob"`
	Phone_Number string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type NurseResponse struct {
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Fullname     string    `json:"fullname"`
	Address      string    `json:"address"`
	Gender       string    `json:"gender"`
	DOB          string    `json:"dob"`
	Phone_Number string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomainRegister(domain nurses.Domain) NurseRegisterRespons {
	return NurseRegisterRespons{
		Message:      "Nurse Registration Success",
		ID:           domain.ID,
		Fullname:     domain.Fullname,
		Username:     domain.Username,
		Address:      domain.Address,
		Gender:       domain.Gender,
		DOB:          domain.DOB,
		Phone_Number: domain.Phone_Number,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDomainAllNurse(domain nurses.Domain) NurseResponse {
	return NurseResponse{
		ID:           domain.ID,
		Fullname:     domain.Fullname,
		Username:     domain.Username,
		Address:      domain.Address,
		Gender:       domain.Gender,
		DOB:          domain.DOB,
		Phone_Number: domain.Phone_Number,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type NurseLoginRespons struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain nurses.Domain) NurseLoginRespons {
	return NurseLoginRespons{
		Message: "Nurse Login Success",
		Token:   domain.Token,
	}
}
func FromDomainUpdateNurse(domain nurses.Domain) NurseRegisterRespons {
	return NurseRegisterRespons{
		Message:   "Update Nurse Profile Success",
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromNurseListDomain(domain []nurses.Domain) []NurseResponse {
	var response []NurseResponse
	for _, value := range domain {
		response = append(response, FromDomainAllNurse(value))
	}
	return response
}
