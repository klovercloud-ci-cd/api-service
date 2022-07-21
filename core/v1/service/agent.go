package service

// Agent operations
type Agent interface {
	Get(companyId string) (httpCode int, body interface{})
	GetK8sObjs(agent, processId string) (httpCode int, body interface{})
	GetPodsByDaemonSet(agent, processId, daemonSetId string) (httpCode int, body interface{})
	GetPodsByDeployment(agent, processId, deploymentId string) (httpCode int, body interface{})
	GetPodsByReplicaSet(agent, processId, replicaSetId string) (httpCode int, body interface{})
	GetPodsByStatefulSet(agent, processId, statefulSetId string) (httpCode int, body interface{})
	GetTerminalByName(name string) (httpCode int, body interface{})
	Store(agent interface{}, name string) (httpCode int, error error)
}
