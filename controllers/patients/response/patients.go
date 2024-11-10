package response

import (
	"Hospital-Management-System/business/patients"
	"time"
)

type PatientRegisterRespons struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Fullname  string    `json:"fullname"`
	Gender    string    `json:"gender"`
	NIK       int       `json:"nik"`
	Address   string    `json:"address"`
	DOB       string    `json:"dob"`
	No_Rm     string    `json:"no_rm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type PatientResponse struct {
	ID        int       `json:"id:"`
	Fullname  string    `json:"fullname"`
	Gender    string    `json:"gender"`
	NIK       int       `json:"nik"`
	Address   string    `json:"address"`
	DOB       string    `json:"dob"`
	No_Rm     string    `json:"no_rm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain patients.Domain) PatientRegisterRespons {
	return PatientRegisterRespons{
		Message:   "Patient Registration Success",
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Gender:    domain.Gender,
		NIK:       domain.NIK,
		Address:   domain.Address,
		DOB:       domain.DOB,
		No_Rm:     domain.No_Rm,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromDomainAllDoctor(domain patients.Domain) PatientResponse {
	return PatientResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Gender:    domain.Gender,
		NIK:       domain.NIK,
		Address:   domain.Address,
		DOB:       domain.DOB,
		No_Rm:     domain.No_Rm,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdateDoctor(domain patients.Domain) PatientRegisterRespons {
	return PatientRegisterRespons{
		Message:   "Update Patient Data Success",
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Gender:    domain.Gender,
		NIK:       domain.NIK,
		Address:   domain.Address,
		DOB:       domain.DOB,
		No_Rm:     domain.No_Rm,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromDoctorListDomain(domain []patients.Domain) []PatientResponse {
	var response []PatientResponse
	for _, value := range domain {
		response = append(response, FromDomainAllDoctor(value))
	}
	return response
}
