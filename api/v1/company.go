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

type companyApi struct {
	companyService service.Company
	jwtService     service.Jwt
}

// Update... Update repositories
// @Summary Update repositories by company id
// @Description updates repositories
// @Tags Company
// @Produce json
// @Param data body interface{} true "List Of Repositories data"
// @Param id path string true "Company id"
// @Param companyUpdateOption query string true "Company Update Option"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories [PUT]
func (c companyApi) UpdateRepositories(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.UPDATE)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Repository not found!")
		}
	}
	companyUpdateOption := context.QueryParam("companyUpdateOption")
	return context.JSON(c.companyService.UpdateRepositories(id, formData, companyUpdateOption))
}

// Save... Save company
// @Summary Save company
// @Description Saves company
// @Tags Company
// @Produce json
// @Param data body interface{} true "Company data"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies [POST]
func (c companyApi) Save(context echo.Context) error {
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.CREATE)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
	}
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	return context.JSON(c.companyService.Store(formData))
}

// Get.. Get RepositoriesDto by company id
// @Summary Get RepositoriesDto by company id
// @Description Gets RepositoriesDto by company id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories [GET]
func (c companyApi) GetRepositoriesById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Repository not found!")
		}
	}
	option := getQueryOption(context)
	return context.JSON(c.companyService.GetRepositoriesById(id, option))
}

// Get.. Get company
// @Summary Get company by id
// @Description Gets company by id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id} [GET]
func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Company not found!")
		}
	}
	option := getQueryOption(context)
	return context.JSON(c.companyService.GetCompaniesById(nil, id, option))
}

// Get... Get companies
// @Summary Get companies
// @Description Gets companies
// @Tags Company
// @Produce json
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query bool false "Loads Repositories"
// @Param loadApplications query bool false "Loads Applications"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies [GET]
func (c companyApi) GetCompanies(context echo.Context) error {
	option := getQueryOption(context)
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
	}
	return context.JSON(c.companyService.GetCompanies(option))
}

//this function is for set all query param
func getQueryOption(context echo.Context) v1.CompanyQueryOption {
	option := v1.CompanyQueryOption{}
	option.Pagination.Page = context.QueryParam("page")
	option.Pagination.Limit = context.QueryParam("limit")
	option.LoadApplications = context.QueryParam("loadApplications")
	option.LoadRepositories = context.QueryParam("loadRepositories")
	return option
}

// NewCompanyApi returns Company type api
func NewCompanyApi(companyService service.Company, jwtService service.Jwt) api.Company {
	return &companyApi{
		companyService: companyService,
		jwtService:     jwtService,
	}
}
