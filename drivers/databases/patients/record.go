package patients


import (
	"Hospital-Management-System/business/patients"
	"time"

	"gorm.io/gorm"
)

type Patients struct {
	gorm.Model
	ID                    int `gorm:"primary_key"`
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

func toDomain(ev Patients) patients.Domain {
	return patients.Domain{
		ID:                    ev.ID,
		MedicalPrescriptionID: ev.MedicalPrescriptionID,
		MedicalRecordID:       ev.MedicalRecordID,
		Fullname:              ev.Fullname,
		Address:               ev.Address,
		Gender:                ev.Gender,
		NIK:                   ev.NIK,
		No_Rm:                 ev.No_Rm,
		DOB:                   ev.DOB,
		CreatedAt:             ev.CreatedAt,
		UpdatedAt:             ev.UpdatedAt,
	}
}

func fromDomain(domain patients.Domain) Patients {
	return Patients{
		ID:                    domain.ID,
		MedicalPrescriptionID: domain.MedicalPrescriptionID,
		MedicalRecordID:       domain.MedicalRecordID,
		Fullname:              domain.Fullname,
		Address:               domain.Address,
		Gender:                domain.Gender,
		NIK:                   domain.NIK,
		No_Rm:                 domain.No_Rm,
		DOB:                   domain.DOB,
		CreatedAt:             domain.CreatedAt,
		UpdatedAt:             domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Patients) patients.Domain {
	return patients.Domain{
		ID:                    ev.ID,
		MedicalPrescriptionID: ev.MedicalPrescriptionID,
		MedicalRecordID:       ev.MedicalRecordID,
		Fullname:              ev.Fullname,
		Address:               ev.Address,
		Gender:                ev.Gender,
		NIK:                   ev.NIK,
		No_Rm:                 ev.No_Rm,
		DOB:                   ev.DOB,
		CreatedAt:             ev.CreatedAt,
		UpdatedAt:             ev.UpdatedAt,
	}
}
func toDomainList(data []Patients) []patients.Domain {
	result := []patients.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}