package api

import "github.com/labstack/echo/v4"

// DaemonSet api operations
type DaemonSet interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
