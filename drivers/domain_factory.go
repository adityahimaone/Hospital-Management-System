package drivers

import (
	adminDomain "Hospital-Management-System/business/admins"
	doctorDomain "Hospital-Management-System/business/doctors"
	nurseDomain "Hospital-Management-System/business/nurses"
	adminDB "Hospital-Management-System/drivers/databases/admins"
	doctorDB "Hospital-Management-System/drivers/databases/doctors"
	nurseDB "Hospital-Management-System/drivers/databases/nurses"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}

func NewDoctorRepository(conn *gorm.DB) doctorDomain.Repository {
	return doctorDB.NewMysqlDoctorRepository(conn)
}

func NewNurseRepository(conn *gorm.DB) nurseDomain.Repository {
	return nurseDB.NewMysqlNurseRepository(conn)
}