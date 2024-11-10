package response

import (
	"Hospital-Management-System/business/doctors"
	"time"
)

type DoctorRegisterRespons struct {
	Message      string    `json:"message"`
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Fullname     string    `json:"fullname"`
	Specialist   string    `json:"specialist"`
	Address      string    `json:"address"`
	DOB          string    `json:"dob"`
	Phone_Number string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type DoctorResponse struct {
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Fullname     string    `json:"fullname"`
	Specialist   string    `json:"specialist"`
	Address      string    `json:"address"`
	DOB          string    `json:"dob"`
	Phone_Number string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomainRegister(domain doctors.Domain) DoctorRegisterRespons {
	return DoctorRegisterRespons{
		Message:    "Doctor Registration Success",
		ID:         domain.ID,
		Fullname:   domain.Fullname,
		Username:   domain.Username,
		Specialist: domain.Specialist,
		Address:    domain.Address,
		DOB:        domain.DOB,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
func FromDomainAllDoctor(domain doctors.Domain) DoctorResponse {
	return DoctorResponse{
		ID:           domain.ID,
		Fullname:     domain.Fullname,
		Username:     domain.Username,
		Specialist:   domain.Specialist,
		Address:      domain.Address,
		DOB:          domain.DOB,
		Phone_Number: domain.Phone_Number,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type DoctorLoginRespons struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain doctors.Domain) DoctorLoginRespons {
	return DoctorLoginRespons{
		Message: "Doctor Login Success",
		Token:   domain.Token,
	}
}
func FromDomainUpdateDoctor(domain doctors.Domain) DoctorRegisterRespons {
	return DoctorRegisterRespons{
		Message:   "Update Doctor Profile Success",
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromDoctorListDomain(domain []doctors.Domain) []DoctorResponse {
	var response []DoctorResponse
	for _, value := range domain {
		response = append(response, FromDomainAllDoctor(value))
	}
	return response
}
