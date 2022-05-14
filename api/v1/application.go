package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type applicationApi struct {
	applicationService service.Company
	jwtService         service.Jwt
}

// GetAll.. Get all applications
// @Summary Get all applications
// @Description Get all applications
// @Tags Application
// @Produce json
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/applications [GET]
func (a applicationApi) GetAll(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.APPLICATION), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	option := getQueryOption(context)
	return context.JSON(a.applicationService.GetAllApplications(companyId, option))
}

// Get.. Get application by appliction id
// @Summary Get application by appliction id
// @Description Get application by appliction id
// @Tags Application
// @Produce json
// @Param id path string true "application id"
// @Param repositoryId query string true "repository id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/applications/{id} [GET]
func (a applicationApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "application Id is required!")
	}
	var companyId string
	repositoryId := context.QueryParam("repositoryId")
	if repositoryId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "repository id is required",
		})
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.APPLICATION), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}

	return context.JSON(a.applicationService.GetApplicationByApplicationId(companyId, repositoryId, id))
}

// Update... Update Application
// @Summary  Update Application
// @Description Update Application by company id and  repository id
// @Tags Application
// @Accept json
// @Produce json
// @Param data body object true "ApplicationWithUpdateOption Data"
// @Param repositoryId query string true "repository Id"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/applications [POST]
func (a applicationApi) Update(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	repoId := context.QueryParam("repositoryId")
	var companyId string
	companyUpdateOption := context.QueryParam("companyUpdateOption")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.APPLICATION), "", string(enums.UPDATE)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	code, err := a.applicationService.UpdateApplication(companyId, repoId, formData, companyUpdateOption)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return common.GenerateSuccessResponse(context, nil, nil, "Applications updated successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
}

// NewApplicationApi returns Application type api
func NewApplicationApi(applicationService service.Company, jwtService service.Jwt) api.Application {
	return &applicationApi{
		applicationService: applicationService,
		jwtService:         jwtService,
	}
}
