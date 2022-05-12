package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type v1BitbucketApi struct {
	bitbucket  service.Bitbucket
	jwtService service.Jwt
}

// UpdateWebhook... Update Webhook
// @Summary Update Webhook to Enable or Disable
// @Description Update Webhook
// @Tags Bitbucket
// @Produce json
// @Param action query string true "action type [enable/disable]"
// @Param companyId query string true "Company Id"
// @Param repoId query string true "Repository Id"
// @Param url query string true "Url"
// @Param webhookId query string true "Webhook Id"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets/webhooks [PATCH]
func (b v1BitbucketApi) UpdateWebhook(context echo.Context) error {
	action := context.QueryParam("action")
	if action == "enable" {
		return b.EnableWebhook(context)
	} else if action == "disable" {
		return b.DisableWebhook(context)
	}
	return common.GenerateErrorResponse(context, nil, "Provide valid action. [enable/disable]")
}

func (v v1BitbucketApi) EnableWebhook(context echo.Context) error {
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
	code, err := v.bitbucket.EnableWebhook(companyId, repoId, userName, repoName)
	if err != nil {
		return err
	}
	if code == 200 {
		return context.JSON(200, "Webhook is enabled")
	} else {
		return context.JSON(code, "Webhook is not enabled")
	}
}

func (v v1BitbucketApi) DisableWebhook(context echo.Context) error {
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
	code, err := v.bitbucket.DisableWebhook(companyId, repoId, userName, repoName, webhookId)
	if err != nil {
		return err
	}
	if code == 200 {
		return context.JSON(200, "Webhook is disabled")
	} else {
		return context.JSON(code, "Webhook is not disabled")
	}
}

// GetCommitByBranch... Get Commit By Branch
// @Summary Get Commit By Branch
// @Description Gets Commit By Branch
// @Tags Bitbucket
// @Produce json
// @Param repoId query string true "Repository Id"
// @Param url query string true "Url"
// @Param branch query string true "Branch"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets/commits [GET]
func (v v1BitbucketApi) GetCommitByBranch(context echo.Context) error {
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
	return context.JSON(v.bitbucket.GetCommitByBranch(repoName, userName, repoId, branch, companyId))
}

// GetBranches... Get Branches
// @Summary Get Branches
// @Description Gets Branches
// @Tags Bitbucket
// @Produce json
// @Param userName query string true "User Name"
// @Param repoId query string true "Repository Id"
// @Param repoName query string true "Repository Name"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets/branches [GET]
func (v v1BitbucketApi) GetBranches(context echo.Context) error {
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
	return context.JSON(v.bitbucket.GetBranches(url, repoId, companyId))
}

// Listen ... Listen Bitbucket Web hook event
// @Summary  Listen Bitbucket Web hook event
// @Description Listens Bitbucket Web hook events. Register this endpoint as Bitbucket web hook endpoint
// @Tags Bitbucket
// @Accept json
// @Produce json
// @Param data body v1.BitbucketWebHookEvent true "GithubWebHookEvent Data"
// @Success 200 {object} common.ResponseDTO{data=string}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets [POST]
func (v v1BitbucketApi) ListenEvent(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	err := v.bitbucket.ListenEvent(formData, companyId)
	if err != nil {
		return err
	}
	return nil
}

// newBitbucketApi returns bitbucket type api
func newBitbucketApi(bitbucket service.Bitbucket, jwtService service.Jwt) api.Git {
	return &v1BitbucketApi{
		bitbucket:  bitbucket,
		jwtService: jwtService,
	}
}
