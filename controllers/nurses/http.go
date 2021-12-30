package nurses

import (
	"net/http"

	"Hospital-Management-System/business/nurses"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/nurses/request"
	response "Hospital-Management-System/controllers/nurses/res"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NurseController struct {
	nurseService nurses.Service
}

func NewControllerDoctor(serv nurses.Service) *NurseController {
	return &NurseController{
		nurseService: serv,
	}
}

func (ctrl *NurseController) Register(c echo.Context) error {

	registerReq := request.Nurses{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.nurseService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *NurseController) AllNurse(c echo.Context) error {

	result, err := ctrl.nurseService.AllNurse()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromNurseListDomain(result))

}

func (ctrl *NurseController) Login(c echo.Context) error {

	loginReq := request.NurseLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.nurseService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainLogin(result))
}

func (ctrl *NurseController) Update(c echo.Context) error {

	updateReq := request.Nurses{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.nurseService.NurseByID(id)
	result, err := ctrl.nurseService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID
	result.Fullname = getData.Fullname

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateNurse(result))

}

func (ctrl *NurseController) NurseByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.nurseService.NurseByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllNurse(result))
}

func (ctrl *NurseController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.nurseService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
