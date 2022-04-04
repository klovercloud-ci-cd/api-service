package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type v1BitbucketApi struct {
	bitbucket service.Bitbucket
}

func (v v1BitbucketApi) GetBranches(context echo.Context) error {
	//TODO implement me
	panic("implement me")
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
func newBitbucketApi(bitbucket service.Bitbucket) api.Git {
	return &v1BitbucketApi{
		bitbucket: bitbucket,
	}
}
