package doctors

import (
	"net/http"

	"Hospital-Management-System/business/doctors"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/doctors/request"
	response "Hospital-Management-System/controllers/doctors/res"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	doctorService doctors.Service
}

func NewControllerDoctor(serv doctors.Service) *DoctorController {
	return &DoctorController{
		doctorService: serv,
	}
}

func (ctrl *DoctorController) Register(c echo.Context) error {

	registerReq := request.Doctors{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.doctorService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}
func (ctrl *DoctorController) AllDoctor(c echo.Context) error {

	result, err := ctrl.doctorService.AllDoctor()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDoctorListDomain(result))

}
func (ctrl *DoctorController) Login(c echo.Context) error {

	loginReq := request.DoctorLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.doctorService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainLogin(result))
}
func (ctrl *DoctorController) Update(c echo.Context) error {

	updateReq := request.Doctors{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.doctorService.DoctorByID(id)
	result, err := ctrl.doctorService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	result.Name = getData.Name

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateDoctor(result))

}
func (ctrl *DoctorController) SellerByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.doctorService.DoctorByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllDoctor(result))
}
func (ctrl *DoctorController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.doctorService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
