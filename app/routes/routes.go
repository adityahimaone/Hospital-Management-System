package routes

import (
	middlewareApp "Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"
	controller "Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/admins"
	"Hospital-Management-System/controllers/doctors"
	"Hospital-Management-System/controllers/nurses"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	doctors.DoctorController
	AdminController admins.AdminController
	nurses.NurseController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// Admins
	e.POST("/api/v1/admins/login", cl.AdminController.Login)

	e.POST("/api/v1/admins/add/doctor", cl.DoctorController.Register, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/doctor/:id", cl.DoctorController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/doctor/:id", cl.DoctorController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/doctor", cl.DoctorController.AllDoctor, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	e.POST("/api/v1/admins/add/nurse", cl.NurseController.Register, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/nurse/:id", cl.NurseController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/nurse/:id", cl.NurseController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/nurse", cl.NurseController.AllNurse, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/nurse/:id", cl.NurseController.NurseByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	// Doctors
	e.POST("/api/v1/doctors/login", cl.DoctorController.Login)
}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}
