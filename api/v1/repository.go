package v1

import (
	"errors"
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

func (r repositoryApi) GetApplicationsById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Company not found!")
		}
	}
	option := getRepositoryQueryOption(context)
	return context.JSON(r.repositoryService.GetApplicationsByCompanyId(id, option))
}
func getRepositoryQueryOption(context echo.Context) v1.RepositoryQueryOption {
	option := v1.RepositoryQueryOption{}
	option.Pagination.Page = context.QueryParam("page")
	option.Pagination.Limit = context.QueryParam("limit")
	option.LoadApplications = context.QueryParam("loadApplications")
	return option
}

func (r repositoryApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
	}
	return context.JSON(r.repositoryService.GetRepositoryByRepositoryId(id))
}

// NewRepositoryApi returns Repository type api
func NewRepositoryApi(repositoryService service.Company, jwtService service.Jwt) api.Repository {
	return &repositoryApi{
		repositoryService: repositoryService,
		jwtService:        jwtService,
	}
}
