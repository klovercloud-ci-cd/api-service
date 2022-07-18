package api

import "github.com/labstack/echo/v4"

// Agent operations
type Agent interface {
	Get(context echo.Context) error
	GetK8sObjs(context echo.Context) error
	GetPodsByCertificate(context echo.Context) error
	GetPodsByClusterRole(context echo.Context) error
	GetPodsByClusterRoleBinding(context echo.Context) error
	GetPodsByConfigMap(context echo.Context) error
	GetPodsByDaemonSet(context echo.Context) error
	GetPodsByDeployment(context echo.Context) error
	GetPodsByIngress(context echo.Context) error
	GetPodsByNamespace(context echo.Context) error
	GetPodsByNetworkPolicy(context echo.Context) error
	GetPodsByNode(context echo.Context) error
	GetPodsByPV(context echo.Context) error
	GetPodsByPVC(context echo.Context) error
	GetPodsByReplicaSet(context echo.Context) error
	GetPodsByRole(context echo.Context) error
	GetPodsByRoleBinding(context echo.Context) error
	GetPodsBySecret(context echo.Context) error
	GetPodsByService(context echo.Context) error
	GetPodsByServiceAccount(context echo.Context) error
	GetPodsByStatefulSet(context echo.Context) error
	GetTerminalByName(context echo.Context) error
	Save(context echo.Context) error
}
