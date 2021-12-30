package patients

import "time"

type Domain struct {
	ID                    int
	MedicalPrescriptionID int
	MedicalRecordID       int
	Fullname              string
	Address               string
	Gender                string
	NIK                   int
	No_Rm                 string
	DOB                   string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type Service interface {
	AllPatient() ([]Domain, error)
	Register(domain *Domain) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	PatientByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllPatient() ([]Domain, error)
	Register(domain *Domain) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	PatientByID(id int) (Domain, error)
	Delete(id int) (string, error)
}
