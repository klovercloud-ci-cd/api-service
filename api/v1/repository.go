package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type repositoryApi struct {
	repositoryService service.Company
	jwtService        service.Jwt
}

// Get.. Get applications by repository id
// @Summary Get applications by repository id
// @Description Get applications by repository id
// @Tags Repository
// @Produce json
// @Param id path string true "repository id"
// @Param status query string false "status"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/repositories/{id}/applications [GET]
func (r repositoryApi) GetApplicationsById(context echo.Context) error {
	var companyId string
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no repositoryId is provided", "Please provide repositoryId")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	option := getRepositoryQueryOption(context)
	status := context.QueryParam("status")
	code, data := r.repositoryService.GetApplicationsByRepositoryId(id, companyId, option, status)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Applications Query by Repository ID Failed", "Operation Failed")
}
func getRepositoryQueryOption(context echo.Context) v1.RepositoryQueryOption {
	option := v1.RepositoryQueryOption{}
	option.Pagination.Page = context.QueryParam("page")
	option.Pagination.Limit = context.QueryParam("limit")
	option.LoadApplications = context.QueryParam("loadApplications")
	return option
}

// Get.. Get repository by repository id
// @Summary Get repository by repository id
// @Description Get repository by repository id
// @Tags Repository
// @Produce json
// @Param id path string true "repository id"
// @Param loadApplications query bool false "Load applications"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/repositories/{id} [GET]
func (r repositoryApi) GetById(context echo.Context) error {
	var companyId string
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no repositoryId is provided", "Please provide repositoryId")
	}
	loadApplications := context.QueryParam("loadApplications")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	code, data := r.repositoryService.GetRepositoryByRepositoryId(id, companyId, loadApplications)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Repository Query by ID Failed", "Operation Failed")
}

// NewRepositoryApi returns Repository type api
func NewRepositoryApi(repositoryService service.Company, jwtService service.Jwt) api.Repository {
	return &repositoryApi{
		repositoryService: repositoryService,
		jwtService:        jwtService,
	}
}
