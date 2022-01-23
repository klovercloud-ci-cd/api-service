package v1

import (
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type processApi struct {
	processService service.Process
	jwtService     service.Jwt
}

// Get... Get Process List or count process
// @Summary Get Process List or count process
// @Description Get Process List or count process
// @Tags Process
// @Produce json
// @Param companyId query string true "Company Id"
// @Param repositoryId query string false "Repository Id"
// @Param appId query string true "App Id"
// @Param operation query string false "Operation[countTodaysProcessByCompanyId]"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/processes [GET]
func (p processApi) GetByCompanyIdAndRepositoryIdAndAppId(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	repositoryId := context.QueryParam("repositoryId")
	appId := context.QueryParam("appId")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if companyId != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Company not found!")
		}
	}
	return context.JSON(p.processService.GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId))
}

// NewProcessApi returns Process type api
func NewProcessApi(processService service.Process, jwtService service.Jwt) api.Process {
	return &processApi{
		processService: processService,
		jwtService:     jwtService,
	}
}
