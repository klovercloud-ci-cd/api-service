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
	LogEventRouter(g.Group("/logs"))
	ProcessEventRouter(g.Group("/processes_events"))
	KubeEventRouter(g.Group("/kube_events"))
	KubeObjectRouter(g.Group("/kube_objects"))
	AgentRouter(g.Group("/agents"))
}

// ProcessEventRouter api/v1/process_events router
func ProcessEventRouter(g *echo.Group) {
	var processEvent api.ProcessEvent
	processEvent = NewProcessEvent(dependency.GetProcessEvent(), dependency.GetV1JwtService())
	g.POST("", processEvent.Save, AuthenticationHandlerForInternalCall)
}

// LogEventRouter api/v1/logs router
func LogEventRouter(g *echo.Group) {
	var logEvent api.LogEvent
	logEvent = NewLogEvent(dependency.GetLogEventService(), dependency.GetV1JwtService())
	g.POST("", logEvent.Save, AuthenticationHandlerForInternalCall)
}

// BitbucketEventRouter api/v1/bitbuckets event router
func BitbucketEventRouter(g *echo.Group) {
	var bitbucket api.Git
	bitbucket = newBitbucketApi(dependency.GetV1BitbucketService(), dependency.GetV1JwtService())
	g.POST("", bitbucket.ListenEvent)
	g.GET("/branches", bitbucket.GetBranches)
	g.GET("/commits", bitbucket.GetCommitByBranch)
}

// GithubEventRouter api/v1/githubs event router
func GithubEventRouter(g *echo.Group) {
	var githubApi api.Git
	githubApi = NewGithubApi(dependency.GetV1GithubService(), dependency.GetV1JwtService())
	g.POST("", githubApi.ListenEvent)
	g.GET("/branches", githubApi.GetBranches)
	g.GET("/commits", githubApi.GetCommitByBranch)
}

// ApplicationRouter api/v1/applications/* router
func ApplicationRouter(g *echo.Group) {
	applicationApi := NewApplicationApi(dependency.GetV1CompanyService(), dependency.GetV1JwtService())
	g.POST("", applicationApi.Update)
	g.GET("/:id", applicationApi.GetById)
	g.GET("", applicationApi.GetAll)
	g.POST("/:id/pipelines", applicationApi.CreatePipeline)
	g.PUT("/:id/pipelines", applicationApi.UpdatePipeline)
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
	g.GET("", pipelineApi.Get)
	g.GET("/:id", pipelineApi.GetByProcessId)
	g.GET("/ws", pipelineApi.GetEvents)
	g.POST("", pipelineApi.Create)
	g.PUT("", pipelineApi.Update)
}

// ProcessRouter api/v1/processes/* router
func ProcessRouter(g *echo.Group) {
	processApi := NewProcessApi(dependency.GetV1ProcessService(), dependency.GetV1JwtService())
	g.GET("", processApi.Get)
	g.GET("/:processId", processApi.GetById)
	g.GET("/:processId/steps/:step/footmarks", processApi.GetFootmarksByProcessIdAndStep)
	g.GET("/:processId/steps/:step/footmarks/:footmark/logs", processApi.GetLogsByProcessIdAndStepAndFootmark)
}

// CompanyRouter api/v1/companies/* router
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(), dependency.GetV1JwtService())
	g.POST("", companyApi.Save)
	g.GET("", companyApi.Get)
	g.GET("/:id", companyApi.GetById)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById)
	g.PUT("/:id/repositories", companyApi.UpdateRepositories)
	g.GET("/:id/applications", companyApi.GetApplicationsByCompanyIdAndRepositoryType)
	g.PATCH("/:id/repositories/:repoId/webhooks", companyApi.UpdateWebhook)
}

// ProcessLifeCycleEventRouter api/v1/process_life_cycle_events/* router
func ProcessLifeCycleEventRouter(g *echo.Group) {
	processLifeCycleEventApi := NewProcessLifeCycleEventApi(dependency.GetV1ProcessLifeCycleEventService(), dependency.GetV1JwtService())
	g.POST("", processLifeCycleEventApi.Save, AuthenticationHandlerForInternalCall)
	g.GET("", processLifeCycleEventApi.Pull, AuthenticationHandlerForInternalCall)
}

// KubeEventRouter api/v1/kube_events/* router
func KubeEventRouter(g *echo.Group) {
	kubeEventApi := NewKubeEventApi(dependency.GetKubeEvent(), dependency.GetV1JwtService())
	g.POST("", kubeEventApi.Save, AuthenticationHandlerForInternalCall)
}

// KubeObjectRouter api/v1/kube_objects/* router
func KubeObjectRouter(g *echo.Group) {
	kubeObjectApi := NewKubeObjectApi(dependency.GetKubeObjectService(), dependency.GetV1JwtService())
	g.GET("", kubeObjectApi.Get, AuthenticationHandlerForInternalCall)
}

// AgentRouter api/v1/agents/* router
func AgentRouter(g *echo.Group) {
	agentApi := NewAgentApi(dependency.GetV1Agent(), dependency.GetV1JwtService())
	g.POST("", agentApi.Save, AuthenticationHandlerForInternalCall)
	g.GET("/:name", agentApi.Get)
}
