package api

import "github.com/labstack/echo/v4"

// KubeObject k8s Object api operations
type KubeObject interface {
	Get(context echo.Context) error
}
