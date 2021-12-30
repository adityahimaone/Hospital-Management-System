package patients

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/patients"

	"gorm.io/gorm"
)

type MysqlPatientRepository struct {
	Conn *gorm.DB
}

func NewMysqlPatientRepository(conn *gorm.DB) patients.Repository {
	return &MysqlPatientRepository{
		Conn: conn,
	}
}
func (rep *MysqlPatientRepository) AllPatient() ([]patients.Domain, error) {

	var patient []Patients

	result := rep.Conn.Find(&patient)

	if result.Error != nil {
		return []patients.Domain{}, result.Error
	}

	return toDomainList(patient), nil

}
func (rep *MysqlPatientRepository) Register(domain *patients.Domain) (patients.Domain, error) {

	user := fromDomain(*domain)

	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return patients.Domain{}, result.Error
	}

	return toDomain(user), nil
}

func (rep *MysqlPatientRepository) Update(patientID int, domain *patients.Domain) (patients.Domain, error) {

	profileUpdate := fromDomain(*domain)

	profileUpdate.ID = patientID

	result := rep.Conn.Where("id = ?", patientID).Updates(&profileUpdate)

	if result.Error != nil {
		return patients.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(profileUpdate), nil
}
func (rep *MysqlPatientRepository) PatientByID(id int) (patients.Domain, error) {

	var patient Patients

	result := rep.Conn.Where("id = ?", id).First(&patient)

	if result.Error != nil {
		return patients.Domain{}, result.Error
	}

	return toDomain(patient), nil
}
func (rep *MysqlPatientRepository) Delete(id int) (string, error) {
	rec := Patients{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Patient has been delete", nil

}
