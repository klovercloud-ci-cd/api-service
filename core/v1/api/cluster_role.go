package api

import "github.com/labstack/echo/v4"

// ClusterRole api operations
type ClusterRole interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
