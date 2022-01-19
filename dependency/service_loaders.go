package dependency

import (
	"github.com/klovercloud-ci-cd/api-service/core/v1/logic"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

// GetV1CompanyService returns Company service
func GetV1CompanyService() service.Company {
	return logic.NewCompanyService(logic.NewHttpClientService())
}

// GetV1ProcessLifeCycleEventService returns ProcessLifeCycleEvent service
func GetV1ProcessLifeCycleEventService() service.ProcessLifeCycleEvent {
	return logic.NewProcessLifeCycleEventService(logic.NewHttpClientService())
}

// GetV1GithubService returns Github service
func GetV1GithubService() service.Github {
	return logic.NewGithubService(logic.NewHttpClientService())
}

// GetV1BitbucketService returns Bitbucket service
func GetV1BitbucketService() service.Bitbucket {
	return logic.NewBitbucketService(logic.NewHttpClientService())
}

// GetV1ProcessService returns Process service
func GetV1ProcessService() service.Process {
	return logic.NewProcessService(logic.NewHttpClientService())
}

// GetV1PipelineService returns Pipeline service
func GetV1PipelineService() service.Pipeline {
	return logic.NewPipelineService(logic.NewHttpClientService(), logic.NewWebsocketClientService())
}

// GetV1JwtService returns Jwt service
func GetV1JwtService() service.Jwt {
	return logic.NewJwtService()
}
