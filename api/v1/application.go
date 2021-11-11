package v1

import (
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type applicationApi struct {
	applicationService service.Company
}

// Update ... Update Application
// @Summary  Update Application
// @Description Update Application by company id and  repository id
// @Tags Application
// @Accept json
// @Produce json
// @Param data body object true "ApplicationWithUpdateOption Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/applications [POST]
func (a applicationApi) Update(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	repoId := context.QueryParam("repository_Id")
	companyId := context.QueryParam("company_Id")
	companyUpdateOption := context.QueryParam("companyUpdateOption")
	return context.JSON(a.applicationService.UpdateApplication(companyId, repoId, formData, companyUpdateOption))
}

// NewApplicationApi returns Application type api
func NewApplicationApi(applicationService service.Company) api.Application {
	return &applicationApi{applicationService: applicationService}
}
