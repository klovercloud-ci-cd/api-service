package api

import "github.com/labstack/echo/v4"

// KubeEvent k8s Event api operations
type KubeEvent interface {
	Save(context echo.Context) error
}
