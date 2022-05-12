package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type v1GithubApi struct {
	github     service.Github
	jwtService service.Jwt
}

// UpdateWebhook... Update Webhook
// @Summary Update Webhook to Enable or Disable
// @Description Update Webhook
// @Tags Github
// @Produce json
// @Param action query string true "action type [enable/disable]"
// @Param companyId query string true "Company Id"
// @Param repoId query string true "Repository Id"
// @Param url query string true "Url"
// @Param webhookId query string true "Webhook Id"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/githubs/webhooks [PATCH]
func (g v1GithubApi) UpdateWebhook(context echo.Context) error {
	action := context.QueryParam("action")
	if action == "enable" {
		return g.EnableWebhook(context)
	} else if action == "disable" {
		return g.DisableWebhook(context)
	}
	return common.GenerateErrorResponse(context, nil, "Provide valid action. [enable/disable]")
}

func (v v1GithubApi) EnableWebhook(context echo.Context) error {
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
	userName := context.QueryParam("userName")
	repoName := context.QueryParam("repoName")
	return context.JSON(v.github.EnableWebhook(companyId, repoId, userName, repoName))
}

func (v v1GithubApi) DisableWebhook(context echo.Context) error {
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
	userName := context.QueryParam("userName")
	repoName := context.QueryParam("repoName")
	webhookId := context.QueryParam("webhookId")
	return context.JSON(v.github.DisableWebhook(companyId, repoId, userName, repoName, webhookId))
}

// GetCommitByBranch... Get commit by branch
// @Summary Get commit by branch
// @Description Gets commit by branch
// @Tags Github
// @Produce json
// @Param userName query string true "User Name"
// @Param branch query string true "branch"
// @Param repoId query string true "Repository Id"
// @Param repoName query string true "Repository Name"
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
	userName := context.QueryParam("userName")
	repoName := context.QueryParam("repoName")
	branch := context.QueryParam("branch")
	return context.JSON(v.github.GetCommitByBranch(userName, repoName, branch, companyId, repoId))
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
	return context.JSON(v.github.GetBranches(url, repoId, companyId))
}

// this is the main function that will be called by the api to listen bitbucket events
func (v v1GithubApi) ListenEvent(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	err := v.github.ListenEvent(formData, companyId)
	if err != nil {
		return err
	}
	return nil
}

// NewGithubApi returns bitbucket type api
func NewGithubApi(github service.Github, jwtService service.Jwt) api.Git {
	return &v1GithubApi{
		github:     github,
		jwtService: jwtService,
	}
}
