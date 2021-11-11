package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type repositoryApi struct {
	repositoryService service.Company
}

func (r repositoryApi) GetApplicationsById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	option := getQueryOption(context)
	return context.JSON(r.repositoryService.GetApplicationsByCompanyId(id, option))
}

func (r repositoryApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}

	return context.JSON(r.repositoryService.GetRepositoryByRepositoryId(id))
}

// NewRepositoryApi returns Repository type api
func NewRepositoryApi(repositoryService service.Company) api.Repository {
	return &repositoryApi{repositoryService: repositoryService}
}
