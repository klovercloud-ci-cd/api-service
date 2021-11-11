package dependency

import (
	"github.com/klovercloud-ci-cd/api-service/core/v1/logic"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

// GetV1CompanyService returns Company services
func GetV1CompanyService() service.Company {
	return logic.NewCompanyService(logic.NewHttpClientService())
}

// GetV1ProcessService returns Process services
func GetV1ProcessService() service.Process {
	return logic.NewProcessService(logic.NewHttpClientService())
}

// GetV1PipelineService returns Pipeline services
func GetV1PipelineService() service.Pipeline {
	return logic.NewPipelineService(logic.NewHttpClientService(), logic.NewWebsocketClientService())
}

// GetV1JwtService returns Jwt services
func GetV1JwtService() service.Jwt {
	return logic.NewJwtService()
}
