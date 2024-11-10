package nurses

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/nurses"

	"gorm.io/gorm"
)

type MysqlNurseRepository struct {
	Conn *gorm.DB
}

func NewMysqlNurseRepository(conn *gorm.DB) nurses.Repository {
	return &MysqlNurseRepository{
		Conn: conn,
	}
}

func (rep *MysqlNurseRepository) Register(domain *nurses.Domain) (nurses.Domain, error) {

	org := fromDomain(*domain)

	result := rep.Conn.Create(&org)

	if result.Error != nil {
		return nurses.Domain{}, result.Error
	}

	return toDomain(org), nil
}

func (rep *MysqlNurseRepository) AllNurse() ([]nurses.Domain, error) {

	var nurse []Nurses

	result := rep.Conn.Find(&nurse)

	if result.Error != nil {
		return []nurses.Domain{}, result.Error
	}

	return toDomainList(nurse), nil

}

func (rep *MysqlNurseRepository) Login(username, password string) (nurses.Domain, error) {
	var org Nurses
	err := rep.Conn.First(&org, "username = ?", username).Error

	if err != nil {
		return nurses.Domain{}, business.ErrEmailorPass
	}

	return toDomain(org), nil
}

func (rep *MysqlNurseRepository) Update(nurseID int, domain *nurses.Domain) (nurses.Domain, error) {

	profileUpdate := fromDomain(*domain)

	profileUpdate.ID = nurseID

	result := rep.Conn.Where("id = ?", nurseID).Updates(&profileUpdate)

	if result.Error != nil {
		return nurses.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(profileUpdate), nil
}

func (rep *MysqlNurseRepository) NurseByID(id int) (nurses.Domain, error) {

	var nurse Nurses

	result := rep.Conn.Where("id = ?", id).First(&nurse)

	if result.Error != nil {
		return nurses.Domain{}, result.Error
	}

	return toDomain(nurse), nil
}

func (rep *MysqlNurseRepository) Delete(id int) (string, error) {
	rec := Nurses{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Nurses has been delete", nil

}