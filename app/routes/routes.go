package routes

import (
	middlewareApp "Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"
	controller "Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/admins"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig

	AdminController admins.AdminController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// Admins

	e.POST("/api/v1/admins/login", cl.AdminController.Login)

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
