package main

import (
	"log"

	_routes "Hospital-Management-System/app/routes"

	_adminService "Hospital-Management-System/business/admins"
	_adminController "Hospital-Management-System/controllers/admins"
	_adminRepo "Hospital-Management-System/drivers/databases/admins"

	_dbDriver "Hospital-Management-System/drivers/mysql"

	_driverFactory "Hospital-Management-System/drivers"

	_middlewares "Hospital-Management-System/app/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_adminRepo.Admins{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminService := _adminService.NewServiceAdmin(adminRepo, 10, &configJWT)
	adminCtrl := _adminController.NewControllerAdmin(adminService)

	routesInit := _routes.ControllerList{
		JWTMiddleware: configJWT.Init(),

		AdminController: *adminCtrl,
	}

	routesInit.RouteRegister(e)
	_middlewares.Logger(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
