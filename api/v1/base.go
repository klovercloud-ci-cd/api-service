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
	CertificateRouter(g.Group("/certificates"))
	ClusterRoleRouter(g.Group("/cluster-roles"))
	ClusterRoleBindingRouter(g.Group("/cluster-role-bindings"))
	ConfigMapRouter(g.Group("/config-maps"))
	DaemonSetRouter(g.Group("/daemon-sets"))
	DeploymentRouter(g.Group("/deployments"))
	IngressRouter(g.Group("/ingresses"))
	NamespaceRouter(g.Group("/namespaces"))
	NetworkPolicyRouter(g.Group("/network-policies"))
	NodeRouter(g.Group("/nodes"))
	PodRouter(g.Group("/pods"))
	PersistentVolumeRouter(g.Group("/persistent-volumes"))
	PersistentVolumeClaimRouter(g.Group("/persistent-volume-claims"))
	ReplicaSetRouter(g.Group("/replica-sets"))
	RoleRouter(g.Group("/roles"))
	RoleBindingRouter(g.Group("/role-bindings"))
	SecretRouter(g.Group("/secrets"))
	ServiceRouter(g.Group("/services"))
	ServiceAccountRouter(g.Group("/service-accounts"))
	StatefulSetRouter(g.Group("/stateful-sets"))
}

// ProcessEventRouter api/v1/process_events router
func ProcessEventRouter(g *echo.Group) {
	var processEvent api.ProcessEvent
	processEvent = NewProcessEvent(dependency.GetProcessEvent(), dependency.GetV1JwtService())
	g.GET("", processEvent.Get)
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
	g.GET("/:processId/process_life_cycle_events", processApi.GetProcessLifeCycleEventByProcessIdAndStepName)
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
	g.PUT("/:id/repositories/:repoId/applications", companyApi.UpdateApplications)
	g.GET("/:id/applications", companyApi.GetApplicationsByCompanyIdAndRepositoryType)
	g.PATCH("/:id/repositories/:repoId/webhooks", companyApi.UpdateWebhook)
}

// ProcessLifeCycleEventRouter api/v1/process_life_cycle_events/* router
func ProcessLifeCycleEventRouter(g *echo.Group) {
	processLifeCycleEventApi := NewProcessLifeCycleEventApi(dependency.GetV1ProcessLifeCycleEventService(), dependency.GetV1JwtService())
	g.POST("", processLifeCycleEventApi.Save, AuthenticationHandlerForInternalCall)
	g.GET("", processLifeCycleEventApi.Pull, AuthenticationHandlerForInternalCall)
	g.PUT("", processLifeCycleEventApi.Update)
}

// KubeEventRouter api/v1/kube_events/* router
func KubeEventRouter(g *echo.Group) {
	kubeEventApi := NewKubeEventApi(dependency.GetKubeEvent(), dependency.GetV1JwtService())
	g.POST("", kubeEventApi.Save, AuthenticationHandlerForInternalCall)
}

// KubeObjectRouter api/v1/kube_objects/* router
func KubeObjectRouter(g *echo.Group) {
	kubeObjectApi := NewKubeObjectApi(dependency.GetKubeObjectService(), dependency.GetV1JwtService())
	g.GET("", kubeObjectApi.Get)
}

// AgentRouter api/v1/agents/* router
func AgentRouter(g *echo.Group) {
	agentApi := NewAgentApi(dependency.GetV1Agent(), dependency.GetV1JwtService())
	g.GET("", agentApi.Get)
	g.GET("/:agent", agentApi.GetByName)
	g.GET("/:agent/k8sobjs", agentApi.GetK8sObjs)
	g.GET("/:agent/daemonSets/:daemonSetId/pods", agentApi.GetPodsByDaemonSet)
	g.GET("/:agent/deployments/:deploymentId/pods", agentApi.GetPodsByDeployment)
	g.GET("/:agent/replicaSets/:replicaSetId/pods", agentApi.GetPodsByReplicaSet)
	g.GET("/:agent/statefulSets/:statefulSetId/pods", agentApi.GetPodsByStatefulSet)
	g.GET("/:name", agentApi.GetTerminalByName)
	g.POST("", agentApi.Save, AuthenticationHandlerForInternalCall)
}

// CertificateRouter api/v1/certificates/* router
func CertificateRouter(g *echo.Group) {
	certificateApi := NewCertificateApi(dependency.GetV1CertificateService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", certificateApi.Get)
	g.GET("/:id", certificateApi.GetByID)
}

// ClusterRoleRouter api/v1/cluster-roles/* router
func ClusterRoleRouter(g *echo.Group) {
	clusterRoleApi := NewClusterRoleApi(dependency.GetV1ClusterRoleService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", clusterRoleApi.Get)
	g.GET("/:id", clusterRoleApi.GetByID)

}

// ClusterRoleBindingRouter api/v1/cluster-role-bindings/* router
func ClusterRoleBindingRouter(g *echo.Group) {
	clusterRoleBindingApi := NewClusterRoleBindingApi(dependency.GetV1ClusterRoleBindingService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", clusterRoleBindingApi.Get)
	g.GET("/:id", clusterRoleBindingApi.GetByID)
}

// ConfigMapRouter api/v1/config-maps/* router
func ConfigMapRouter(g *echo.Group) {
	configMapApi := NewConfigMapApi(dependency.GetV1ConfigMapService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", configMapApi.Get)
	g.GET("/:id", configMapApi.GetByID)
}

// DaemonSetRouter api/v1/daemon-sets/* router
func DaemonSetRouter(g *echo.Group) {
	daemonSetApi := NewDaemonSetApi(dependency.GetV1DaemonSetService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", daemonSetApi.Get)
	g.GET("/:id", daemonSetApi.GetByID)
}

// DeploymentRouter api/v1/deployments/* router
func DeploymentRouter(g *echo.Group) {
	deploymentApi := NewDeploymentApi(dependency.GetV1DeploymentService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", deploymentApi.Get)
	g.GET("/:id", deploymentApi.GetByID)
}

// IngressRouter api/v1/ingresses/* router
func IngressRouter(g *echo.Group) {
	ingressApi := NewIngressApi(dependency.GetV1IngressService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", ingressApi.Get)
	g.GET("/:id", ingressApi.GetByID)
}

// NamespaceRouter api/v1/namespaces/* router
func NamespaceRouter(g *echo.Group) {
	namespaceApi := NewNamespaceApi(dependency.GetV1NamespaceService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", namespaceApi.Get)
	g.GET("/:id", namespaceApi.GetByID)
}

// NetworkPolicyRouter api/v1/network-policies/* router
func NetworkPolicyRouter(g *echo.Group) {
	networkPolicyApi := NewNetworkPolicyApi(dependency.GetV1NetworkPolicyService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", networkPolicyApi.Get)
	g.GET("/:id", networkPolicyApi.GetByID)
}

// NodeRouter api/v1/nodes/* router
func NodeRouter(g *echo.Group) {
	nodeApi := NewNodeApi(dependency.GetV1NodeService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", nodeApi.Get)
	g.GET("/:id", nodeApi.GetByID)
}

// PodRouter api/v1/pods/* router
func PodRouter(g *echo.Group) {
	podApi := NewPodApi(dependency.GetV1PodService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", podApi.Get)
	g.GET("/:id", podApi.GetByID)
}

// PersistentVolumeRouter api/v1/persistent-volumes/* router
func PersistentVolumeRouter(g *echo.Group) {
	persistentVolumeApi := NewPersistentVolumeApi(dependency.GetV1PersistentVolumeService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", persistentVolumeApi.Get)
	g.GET("/:id", persistentVolumeApi.GetByID)
}

// PersistentVolumeClaimRouter api/v1/persistent-volume-claims/* router
func PersistentVolumeClaimRouter(g *echo.Group) {
	persistentVolumeClaimApi := NewPersistentVolumeClaimApi(dependency.GetV1PersistentVolumeClaimService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", persistentVolumeClaimApi.Get)
	g.GET("/:id", persistentVolumeClaimApi.GetByID)
}

// ReplicaSetRouter api/v1/replica-sets/* router
func ReplicaSetRouter(g *echo.Group) {
	replicaSetApi := NewReplicaSetApi(dependency.GetV1ReplicaSetService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", replicaSetApi.Get)
	g.GET("/:id", replicaSetApi.GetByID)
}

// RoleRouter api/v1/roles/* router
func RoleRouter(g *echo.Group) {
	roleApi := NewRoleApi(dependency.GetV1RoleService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", roleApi.Get)
	g.GET("/:id", roleApi.GetByID)
}

// RoleBindingRouter api/v1/role-bindings/* router
func RoleBindingRouter(g *echo.Group) {
	roleBindingApi := NewRoleBindingApi(dependency.GetV1RoleBindingService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", roleBindingApi.Get)
	g.GET("/:id", roleBindingApi.GetByID)
}

// SecretRouter api/v1/secrets/* router
func SecretRouter(g *echo.Group) {
	secretApi := NewSecretApi(dependency.GetV1SecretService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", secretApi.Get)
	g.GET("/:id", secretApi.GetByID)
}

// ServiceRouter api/v1/services/* router
func ServiceRouter(g *echo.Group) {
	serviceApi := NewServiceApi(dependency.GetV1ServiceService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", serviceApi.Get)
	g.GET("/:id", serviceApi.GetByID)
}

// ServiceAccountRouter api/v1/service-accounts/* router
func ServiceAccountRouter(g *echo.Group) {
	serviceAccountApi := NewServiceAccountApi(dependency.GetV1ServiceAccountService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", serviceAccountApi.Get)
	g.GET("/:id", serviceAccountApi.GetByID)
}

// StatefulSetRouter api/v1/stateful-sets/* router
func StatefulSetRouter(g *echo.Group) {
	statefulSetApi := NewStatefulSetApi(dependency.GetV1StatefulSetService(), dependency.GetV1JwtService(), dependency.GetV1ProcessService())
	g.GET("", statefulSetApi.Get)
	g.GET("/:id", statefulSetApi.GetByID)
}
