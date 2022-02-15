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
	repoId := context.QueryParam("repositoryId")
	companyId := context.QueryParam("companyId")
	companyUpdateOption := context.QueryParam("companyUpdateOption")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, a.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.APPLICATION), "", string(enums.UPDATE)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if companyId != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Company not found!")
		}
	}
	httpCode, err := a.applicationService.UpdateApplication(companyId, repoId, formData, companyUpdateOption)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if httpCode == 200 || httpCode == 201 {
		return context.JSON(200, common.ResponseDTO{
			Metadata: nil,
			Data:     nil,
			Status:   "success",
			Message:  "Repository updated successfully!",
		})
	} else {
		return context.JSON(httpCode, common.ResponseDTO{
			Metadata: nil,
			Data:     nil,
			Status:   "error",
			Message:  "Repository not updated!",
		})
	}
}

// NewApplicationApi returns Application type api
func NewApplicationApi(applicationService service.Company, jwtService service.Jwt) api.Application {
	return &applicationApi{
		applicationService: applicationService,
		jwtService:         jwtService,
	}
}
