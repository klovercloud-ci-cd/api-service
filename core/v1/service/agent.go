package service

// Agent operations
type Agent interface {
	Get(companyId string) (httpCode int, body interface{})
	GetK8sObjs(agent, processId string) (httpCode int, body interface{})
	GetPodsByCertificate(agent, processId, certificateId string) (httpCode int, body interface{})
	GetPodsByClusterRole(agent, processId, clusterRoleId string) (httpCode int, body interface{})
	GetPodsByClusterRoleBinding(agent, processId, clusterRoleBindingId string) (httpCode int, body interface{})
	GetPodsByConfigMap(agent, processId, configMapId string) (httpCode int, body interface{})
	GetPodsByDaemonSet(agent, processId, daemonSetId string) (httpCode int, body interface{})
	GetPodsByDeployment(agent, processId, deploymentId string) (httpCode int, body interface{})
	GetPodsByIngress(agent, processId, ingressId string) (httpCode int, body interface{})
	GetPodsByNamespace(agent, processId, namespaceId string) (httpCode int, body interface{})
	GetPodsByNetworkPolicy(agent, processId, networkPolicyId string) (httpCode int, body interface{})
	GetPodsByNode(agent, processId, nodeId string) (httpCode int, body interface{})
	GetPodsByPV(agent, processId, PVId string) (httpCode int, body interface{})
	GetPodsByPVC(agent, processId, PVCId string) (httpCode int, body interface{})
	GetPodsByReplicaSet(agent, processId, replicaSetId string) (httpCode int, body interface{})
	GetPodsByRole(agent, processId, roleId string) (httpCode int, body interface{})
	GetPodsByRoleBinding(agent, processId, roleBindingId string) (httpCode int, body interface{})
	GetPodsBySecret(agent, processId, secretId string) (httpCode int, body interface{})
	GetPodsByService(agent, processId, serviceId string) (httpCode int, body interface{})
	GetPodsByServiceAccount(agent, processId, serviceAccountId string) (httpCode int, body interface{})
	GetPodsByStatefulSet(agent, processId, statefulSetId string) (httpCode int, body interface{})
	GetTerminalByName(name string) (httpCode int, body interface{})
	Store(agent interface{}, name string) (httpCode int, error error)
}
