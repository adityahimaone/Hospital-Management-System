package doctors

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/doctors"

	"gorm.io/gorm"
)

type MysqlDoctorRepository struct {
	Conn *gorm.DB
}

func NewMysqlDoctorRepository(conn *gorm.DB) doctors.Repository {
	return &MysqlDoctorRepository{
		Conn: conn,
	}
}

func (rep *MysqlDoctorRepository) Register(domain *doctors.Domain) (doctors.Domain, error) {

	org := fromDomain(*domain)

	result := rep.Conn.Create(&org)

	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return toDomain(org), nil
}
func (rep *MysqlDoctorRepository) AllDoctor() ([]doctors.Domain, error) {

	var doctor []Doctors

	result := rep.Conn.Find(&doctor)

	if result.Error != nil {
		return []doctors.Domain{}, result.Error
	}

	return toDomainList(doctor), nil

}
func (rep *MysqlDoctorRepository) Login(username, password string) (doctors.Domain, error) {
	var org Doctors
	err := rep.Conn.First(&org, "username = ?", username).Error

	if err != nil {
		return doctors.Domain{}, business.ErrEmailorPass
	}

	return toDomain(org), nil
}
func (rep *MysqlDoctorRepository) Update(docID int, domain *doctors.Domain) (doctors.Domain, error) {

	profileUpdate := fromDomain(*domain)

	profileUpdate.ID = docID

	result := rep.Conn.Where("id = ?", docID).Updates(&profileUpdate)

	if result.Error != nil {
		return doctors.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(profileUpdate), nil
}
func (rep *MysqlDoctorRepository) DoctorByID(id int) (doctors.Domain, error) {

	var doctor Doctors

	result := rep.Conn.Where("id = ?", id).First(&doctor)

	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return toDomain(doctor), nil
}
func (rep *MysqlDoctorRepository) Delete(id int) (string, error) {
	rec := Doctors{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Doctor has been delete", nil

}
