package v1

import (
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/dependency"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	CompanyRouter(g.Group("/companies"))
	ProcessRouter(g.Group("/processes"))
	PipelineRouter(g.Group("/pipelines"))
	RepositoryRouter(g.Group("/repositories"))
	ApplicationRouter(g.Group("/applications"))
	GithubEventRouter(g.Group("/githubs"))
	BitbucketEventRouter(g.Group("/bitbuckets"))
	ProcessLifeCycleEventRouter(g.Group("/process_life_cycle_events"))
}

// BitbucketEventRouter api/v1/bitbuckets event router
func BitbucketEventRouter(g *echo.Group) {
	var bitbucket api.Git
	bitbucket = newBitbucketApi(dependency.GetV1BitbucketService())
	g.POST("", bitbucket.ListenEvent)
}

// GithubEventRouter api/v1/githubs/* router
func GithubEventRouter(g *echo.Group) {
	var githubApi api.Git
	githubApi = NewGithubApi(dependency.GetV1GithubService())
	g.POST("", githubApi.ListenEvent)
}

// ApplicationRouter api/v1/applications/* router
func ApplicationRouter(g *echo.Group) {
	applicationApi := NewApplicationApi(dependency.GetV1CompanyService())
	g.POST("", applicationApi.Update)
}

// RepositoryRouter api/v1/repositories/* router
func RepositoryRouter(g *echo.Group) {
	repositoryApi := NewRepositoryApi(dependency.GetV1CompanyService())
	g.GET("/:id", repositoryApi.GetById)
	g.GET("/:id/applications", repositoryApi.GetApplicationsById)
}

// PipelineRouter api/v1/pipelines/* router
func PipelineRouter(g *echo.Group) {
	pipelineApi := NewPipelineApi(dependency.GetV1PipelineService())
	g.GET("/:id", pipelineApi.GetLogs)
	g.GET("/ws", pipelineApi.GetEvents)
}

// ProcessRouter api/v1/processes/* router
func ProcessRouter(g *echo.Group) {
	processApi := NewProcessApi(dependency.GetV1ProcessService())
	g.GET("", processApi.GetByCompanyIdAndRepositoryIdAndAppId)

}

// CompanyRouter api/v1/companies/* router
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(),dependency.GetV1JwtService())
	g.POST("", companyApi.Save,AuthenticationHandler)
	g.GET("", companyApi.GetCompanies)
	g.GET("/:id", companyApi.GetById)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById)
	g.PUT("/:id/repositories", companyApi.UpdateRepositories)
}


// ProcessLifeCycleEventRouter api/v1/process_life_cycle_events/* router
func ProcessLifeCycleEventRouter(g *echo.Group) {
	processLifeCycleEventApi := NewProcessLifeCycleEventApi(dependency.GetV1ProcessLifeCycleEventService())
	g.POST("", processLifeCycleEventApi.Save)
	g.GET("", processLifeCycleEventApi.Pull)
}
