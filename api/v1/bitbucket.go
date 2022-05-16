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
	url := context.QueryParam("url")
	branch := context.QueryParam("branch")
	return context.JSON(v.bitbucket.GetCommitByBranch(url, repoId, branch, companyId))
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

// this is the main function that will be called by the api to listen bitbucket events
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
