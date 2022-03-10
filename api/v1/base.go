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
	applicationApi := NewApplicationApi(dependency.GetV1CompanyService(), dependency.GetV1JwtService())
	g.POST("", applicationApi.Update)
	g.GET("/:id", applicationApi.GetById)
	g.GET("", applicationApi.GetAll)
}

// RepositoryRouter api/v1/repositories/* router
func RepositoryRouter(g *echo.Group) {
	repositoryApi := NewRepositoryApi(dependency.GetV1CompanyService(), dependency.GetV1JwtService())
	g.GET("/:id", repositoryApi.GetById)
	g.GET("/:id/applications", repositoryApi.GetApplicationsById)
}

// PipelineRouter api/v1/pipelines/* router
func PipelineRouter(g *echo.Group) {
	pipelineApi := NewPipelineApi(dependency.GetV1PipelineService(), dependency.GetV1JwtService())
	g.GET("/:id", pipelineApi.GetLogs)
	g.GET("/ws", pipelineApi.GetEvents)
}

// ProcessRouter api/v1/processes/* router
func ProcessRouter(g *echo.Group) {
	processApi := NewProcessApi(dependency.GetV1ProcessService(), dependency.GetV1JwtService())
	g.GET("", processApi.Get)

}

// CompanyRouter api/v1/companies/* router
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(), dependency.GetV1JwtService())
	g.POST("", companyApi.Save, AuthenticationHandler)
	g.GET("", companyApi.Get)
	g.GET("/:id", companyApi.GetById)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById)
	g.PUT("/:id/repositories", companyApi.UpdateRepositories)
	g.GET("/:id/applications", companyApi.GetApplicationsByCompanyIdAndRepositoryType)
}

// ProcessLifeCycleEventRouter api/v1/process_life_cycle_events/* router
func ProcessLifeCycleEventRouter(g *echo.Group) {
	processLifeCycleEventApi := NewProcessLifeCycleEventApi(dependency.GetV1ProcessLifeCycleEventService(), dependency.GetV1JwtService())
	g.POST("", processLifeCycleEventApi.Save, AuthenticationHandlerForInternalCall)
	g.GET("", processLifeCycleEventApi.Pull, AuthenticationHandlerForInternalCall)
}
