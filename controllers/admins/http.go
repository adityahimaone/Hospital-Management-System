package admins

import (
	"Hospital-Management-System/business/admins"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/admins/request"
	"Hospital-Management-System/controllers/admins/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminService admins.Service
}

func NewControllerAdmin(serv admins.Service) *AdminController {
	return &AdminController{
		adminService: serv,
	}
}

func (ctrl *AdminController) Login(c echo.Context) error {

	loginReq := request.AdminLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.adminService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, res.FromDomainLogin(result))
}
