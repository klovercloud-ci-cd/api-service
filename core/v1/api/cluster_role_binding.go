package api

import "github.com/labstack/echo/v4"

// ClusterRoleBinding api operations
type ClusterRoleBinding interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
