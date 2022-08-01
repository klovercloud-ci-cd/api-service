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

// GetLogEventService returns LogEvent service
func GetLogEventService() service.LogEvent {
	return logic.NewLogEventService(logic.NewHttpClientService())
}

// GetKubeObjectService returns KubeObject service
func GetKubeObjectService() service.KubeObject {
	return logic.NewKubeObjectService(logic.NewHttpClientService())
}

// GetProcessEvent returns ProcessEvent service
func GetProcessEvent() service.ProcessEvent {
	return logic.NewProcessEvent(logic.NewHttpClientService())
}

// GetKubeEvent returns KubeEvent service
func GetKubeEvent() service.KubeEvent {
	return logic.NewKubeEventEventService(logic.NewHttpClientService())
}

// GetV1Agent returns Agent service
func GetV1Agent() service.Agent {
	return logic.NewAgentService(logic.NewHttpClientService())
}

// GetV1CertificateService returns service.Certificate
func GetV1CertificateService() service.Certificate {
	return logic.NewCertificateService(logic.NewHttpClientService())
}

// GetV1ClusterRoleService returns service.ClusterRole
func GetV1ClusterRoleService() service.ClusterRole {
	return logic.NewClusterRoleService(logic.NewHttpClientService())
}

// GetV1ClusterRoleBindingService returns service.ClusterRoleBinding
func GetV1ClusterRoleBindingService() service.ClusterRoleBinding {
	return logic.NewClusterRoleBindingService(logic.NewHttpClientService())
}

// GetV1DaemonSetService returns service.DaemonSet
func GetV1DaemonSetService() service.DaemonSet {
	return logic.NewDaemonSetService(logic.NewHttpClientService())
}

// GetV1DeploymentService returns service.Deployment
func GetV1DeploymentService() service.Deployment {
	return logic.NewDeploymentService(logic.NewHttpClientService())
}

// GetV1IngressService returns service.Ingress
func GetV1IngressService() service.Ingress {
	return logic.NewIngressService(logic.NewHttpClientService())
}

// GetV1NamespaceService returns service.Namespace
func GetV1NamespaceService() service.Namespace {
	return logic.NewNamespaceService(logic.NewHttpClientService())
}

// GetV1NetworkPolicyService returns service.NetworkPolicy
func GetV1NetworkPolicyService() service.NetworkPolicy {
	return logic.NewNetworkPolicyService(logic.NewHttpClientService())
}

// GetV1NodeService returns service.Node
func GetV1NodeService() service.Node {
	return logic.NewNodeService(logic.NewHttpClientService())
}

// GetV1PodService returns service.Pod
func GetV1PodService() service.Pod {
	return logic.NewPodService(logic.NewHttpClientService())
}

// GetV1ConfigMapService returns service.ConfigMap
func GetV1ConfigMapService() service.ConfigMap {
	return logic.NewConfigMapService(logic.NewHttpClientService())
}

// GetV1PersistentVolumeService returns service.PersistentVolume
func GetV1PersistentVolumeService() service.PersistentVolume {
	return logic.NewPersistentVolumeService(logic.NewHttpClientService())
}

// GetV1PersistentVolumeClaimService returns service.PersistentVolumeClaim
func GetV1PersistentVolumeClaimService() service.PersistentVolumeClaim {
	return logic.NewPersistentVolumeClaimService(logic.NewHttpClientService())
}

// GetV1ReplicaSetService returns service.ReplicaSet
func GetV1ReplicaSetService() service.ReplicaSet {
	return logic.NewReplicaSetService(logic.NewHttpClientService())
}

// GetV1RoleService returns service.Role
func GetV1RoleService() service.Role {
	return logic.NewRoleService(logic.NewHttpClientService())
}

// GetV1RoleBindingService returns service.RoleBinding
func GetV1RoleBindingService() service.RoleBinding {
	return logic.NewRoleBindingService(logic.NewHttpClientService())
}

// GetV1SecretService returns service.Secret
func GetV1SecretService() service.Secret {
	return logic.NewSecretService(logic.NewHttpClientService())
}

// GetV1ServiceService returns service.Service
func GetV1ServiceService() service.Service {
	return logic.NewServiceService(logic.NewHttpClientService())
}

// GetV1ServiceAccountService returns service.ServiceAccount
func GetV1ServiceAccountService() service.ServiceAccount {
	return logic.NewServiceAccountService(logic.NewHttpClientService())
}

// GetV1StatefulSetService returns service.StatefulSet
func GetV1StatefulSetService() service.StatefulSet {
	return logic.NewStatefulSetService(logic.NewHttpClientService())
}
