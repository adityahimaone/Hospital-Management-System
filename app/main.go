package main

import (
	"log"

	_routes "Hospital-Management-System/app/routes"

	_adminService "Hospital-Management-System/business/admins"
	_adminController "Hospital-Management-System/controllers/admins"
	_adminRepo "Hospital-Management-System/drivers/databases/admins"

	_doctorService "Hospital-Management-System/business/doctors"
	_doctorController "Hospital-Management-System/controllers/doctors"
	_doctorRepo "Hospital-Management-System/drivers/databases/doctors"
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
		&_doctorRepo.Doctors{},
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

	doctorRepo := _driverFactory.NewDoctorRepository(db)
	doctorService := _doctorService.NewServiceDoctor(doctorRepo, 10, &configJWT)
	doctorCtrl := _doctorController.NewControllerDoctor(doctorService)

	routesInit := _routes.ControllerList{
		JWTMiddleware: configJWT.Init(),

		AdminController:  *adminCtrl,
		DoctorController: *doctorCtrl,
	}

	routesInit.RouteRegister(e)
	_middlewares.Logger(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
