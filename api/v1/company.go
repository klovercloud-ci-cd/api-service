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

type companyApi struct {
	companyService service.Company
	jwtService     service.Jwt
}

// Get.. Get Applications by company id and repository type
// @Summary Get Applications by company id and repository type
// @Description Gets RApplications by company id and repository type
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Param repository_type path string true "Repository Type"
// @Param status query string false "status"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query bool false "Loads Repositories"
// @Param loadApplications query bool false "Loads Applications"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/applications [GET]
func (c companyApi) GetApplicationsByCompanyIdAndRepositoryType(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
		}
	}
	repositoryType := context.QueryParam("repository_type")
	option := getQueryOption(context)
	status := context.QueryParam("status")
	code, data := c.companyService.GetApplicationsByCompanyIdAndRepositoryType(id, repositoryType, option, status)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Applications Query Failed", "Operation Failed")
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
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.UPDATE)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Company not found!")
		}
	}
	companyUpdateOption := context.QueryParam("companyUpdateOption")
	code, err := c.companyService.UpdateRepositories(id, formData, companyUpdateOption)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return context.JSON(code, "Repositories Updated Successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
}

// Update... Update repositories
// @Summary Update repositories by company id
// @Description updates repositories
// @Tags Company
// @Produce json
// @Param data body v1.RepositoriesDto true "RepositoriesDto data"
// @Param id path string true "Company id"
// @Param repoId path string true "Repository id"
// @Param companyUpdateOption query string true "Company Update Option"
// @Param validTill query string false "secret validity date in UTC"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories/{repoId}/applications [PUT]
func (c companyApi) UpdateApplications(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company Id is not provided", "Please provide Company Id")
	}
	repoId := context.Param("repoId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository Id is not provided", "Please provide repository id")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.UPDATE)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Company not found!")
		}
	}
	companyUpdateOption := context.QueryParam("companyUpdateOption")
	validity := context.QueryParam("validTill")
	code, err := c.companyService.UpdateApplications(id, repoId, formData, companyUpdateOption, validity)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return context.JSON(code, "Applications Updated Successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
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
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.CREATE)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	code, err := c.companyService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return context.JSON(code, "Company Created Successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
}

// Get.. Get RepositoriesDto by company id
// @Summary Get RepositoriesDto by company id
// @Description Gets RepositoriesDto by company id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query bool false "Loads Repositories"
// @Param loadApplications query bool false "Loads Applications"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories [GET]
func (c companyApi) GetRepositoriesById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return context.JSON(404, "Repository not found!")
		}
	}
	option := getQueryOption(context)
	code, data := c.companyService.GetRepositoriesById(id, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Repositories Query Failed", "Operation Failed")
}

// Get.. Get company
// @Summary Get company by id
// @Description Gets company by id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query bool false "Loads Repositories"
// @Param loadApplications query bool false "Loads Applications"
// @Param action query string false "action [dashboard_data]"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id} [GET]
func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return common.GenerateErrorResponse(context, "[ERROR] company not found", "Please provide valid company id")
		}
	}
	option := getQueryOption(context)
	action := context.QueryParam("action")
	code, data := c.companyService.GetById(nil, id, action, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Company Query by ID Failed", "Operation Failed")
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
// @Param loadToken query bool false "Loads Token"
// @Param status query string false "status"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies [GET]
func (c companyApi) Get(context echo.Context) error {
	option := getQueryOption(context)
	status := context.QueryParam("status")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	code, data := c.companyService.Get(option, status)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Companies Query Failed", "Operation Failed")
}

// UpdateWebhook... Update Webhook
// @Summary Update Webhook to Enable or Disable
// @Description Update Webhook
// @Tags Github
// @Produce json
// @Param action query string true "action type [enable/disable]"
// @Param repoType query string true "Repository type [github/bitbucket]"
// @Param id path string true "Company id"
// @Param repoId path string true "Repository id"
// @Param url query string true "Url"
// @Param webhookId query string false "Webhook Id to disable webhook"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories/{repoId}/webhooks [PATCH]
func (c companyApi) UpdateWebhook(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, c.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.COMPANY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if id != userResourcePermission.Metadata.CompanyId {
			return common.GenerateErrorResponse(context, "[ERROR] company not found", "Please provide valid company id")
		}
	}
	repoId := context.Param("repoId")
	url := context.QueryParam("url")
	webhookId := context.QueryParam("webhookId")
	action := context.QueryParam("action")
	repoType := context.QueryParam("repoType")
	appId := context.QueryParam("appId")
	_, err := c.companyService.UpdateWebhook(id, repoId, appId, url, webhookId, action, repoType)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR] webhook update failed", err.Error())
	}
	return common.GenerateSuccessResponse(context, "[SUCCESS]: webhook update successful", nil, "Webhook Updated!")
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
