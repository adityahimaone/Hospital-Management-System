package patients

import (
	"net/http"

	"Hospital-Management-System/business/patients"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/patients/request"
	response "Hospital-Management-System/controllers/patients/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatientController struct {
	doctorService patients.Service
}

func NewControllerPatient(serv patients.Service) *PatientController {
	return &PatientController{
		doctorService: serv,
	}
}

func (ctrl *PatientController) Register(c echo.Context) error {

	registerReq := request.Patients{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.doctorService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}
func (ctrl *PatientController) AllPatient(c echo.Context) error {

	result, err := ctrl.doctorService.AllPatient()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDoctorListDomain(result))

}

func (ctrl *PatientController) Update(c echo.Context) error {

	updateReq := request.Patients{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.doctorService.PatientByID(id)
	result, err := ctrl.doctorService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateDoctor(result))

}
func (ctrl *PatientController) SellerByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.doctorService.PatientByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllDoctor(result))
}
func (ctrl *PatientController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.doctorService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
