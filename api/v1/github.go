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

type v1GithubApi struct {
	github     service.Github
	jwtService service.Jwt
}

// GetCommitByBranch... Get commit by branch
// @Summary Get commit by branch
// @Description Gets commit by branch
// @Tags Github
// @Produce json
// @Param url query string true "Url"
// @Param branch query string true "branch"
// @Param repoId query string true "Repository Id"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/githubs/commits [GET]
func (v v1GithubApi) GetCommitByBranch(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, v.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	repoId := context.QueryParam("repoId")
	url := context.QueryParam("url")
	branch := context.QueryParam("branch")
	option := getCommitsPaginationOption(context)
	code, data := v.github.GetCommitByBranch(url, repoId, branch, companyId, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Commit Query by Branch Failed", "Operation Failed")
}

func getCommitsPaginationOption(context echo.Context) v1.Pagination {
	var option v1.Pagination
	option.Page = context.QueryParam("page")
	option.Limit = context.QueryParam("limit")
	return option
}

// GetBranches... Get Branches
// @Summary Get Branches
// @Description Gets Branches
// @Tags Github
// @Produce json
// @Param repoId query string true "Repository Id"
// @Param url query string true "Url"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/githubs/branches [GET]
func (v v1GithubApi) GetBranches(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, v.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	repoId := context.QueryParam("repoId")
	url := context.QueryParam("url")
	code, data := v.github.GetBranches(url, repoId, companyId)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Branches Query Failed", "Operation Failed")
}

// this is the main function that will be called by the api to listen bitbucket events
func (v v1GithubApi) ListenEvent(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	companyId := context.QueryParam("companyId")
	bearerToken := context.Request().Header.Get("Authorization")
	userType := enums.CLIENT
	if bearerToken != "" {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, v.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
		userType = enums.REGULAR
	}

	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no company id is provided", "Please provide company id")
	}
	appId := context.QueryParam("appId")
	if appId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no app id is provided", "Please provide app id")
	}
	appSecret := context.QueryParam("appSecret")
	if appSecret == "" && userType == enums.CLIENT {
		return common.GenerateErrorResponse(context, "[ERROR] no application secret is provided", "Please provide app id")
	}
	code, err := v.github.ListenEvent(formData, companyId, appId, appSecret, string(userType))
	if err != nil {
		return common.GenerateGenericResponse(context, code, err.Error(), "Operation Failed")
	}
	return common.GenerateGenericResponse(context, code, nil, "Operation Failed")
}

// NewGithubApi returns bitbucket type api
func NewGithubApi(github service.Github, jwtService service.Jwt) api.Git {
	return &v1GithubApi{
		github:     github,
		jwtService: jwtService,
	}
}
