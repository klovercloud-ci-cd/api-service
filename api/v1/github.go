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

// GetBranches... Get Branches
// @Summary Get Branches
// @Description Gets Branches
// @Tags Github
// @Produce json
// @Param userName query string true "User Name"
// @Param repoId query string true "Repository Id"
// @Param repoName query string true "Repository Name"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/githubs [GET]
func (v v1GithubApi) GetBranches(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, v.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.REPOSITORY), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	repoId := context.QueryParam("repoId")
	userName := context.QueryParam("userName")
	repoName := context.QueryParam("repoName")
	return context.JSON(v.github.GetBranches(repoName, userName, repoId, companyId))
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
