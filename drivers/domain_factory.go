package drivers

import (
	adminDomain "Hospital-Management-System/business/admins"
	doctorDomain "Hospital-Management-System/business/doctors"
	adminDB "Hospital-Management-System/drivers/databases/admins"
	doctorDB "Hospital-Management-System/drivers/databases/doctors"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
func NewDoctorRepository(conn *gorm.DB) doctorDomain.Repository {
	return doctorDB.NewMysqlDoctorRepository(conn)
}
