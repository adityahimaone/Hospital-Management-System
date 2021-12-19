package drivers

import (
	adminDomain "Hospital-Management-System/business/admins"

	adminDB "Hospital-Management-System/drivers/databases/admins"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
