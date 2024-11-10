package drivers

import (
	adminDomain "Hospital-Management-System/business/admins"
	doctorDomain "Hospital-Management-System/business/doctors"
	patientDomain "Hospital-Management-System/business/patients"

	adminDB "Hospital-Management-System/drivers/databases/admins"
	doctorDB "Hospital-Management-System/drivers/databases/doctors"
	patientDB "Hospital-Management-System/drivers/databases/patients"

	nurseDomain "Hospital-Management-System/business/nurses"

	nurseDB "Hospital-Management-System/drivers/databases/nurses"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}

func NewDoctorRepository(conn *gorm.DB) doctorDomain.Repository {
	return doctorDB.NewMysqlDoctorRepository(conn)
}

func NewPatientRepository(conn *gorm.DB) patientDomain.Repository {
	return patientDB.NewMysqlPatientRepository(conn)
}

func NewNurseRepository(conn *gorm.DB) nurseDomain.Repository {
	return nurseDB.NewMysqlNurseRepository(conn)
}
