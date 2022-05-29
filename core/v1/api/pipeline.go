package api

import "github.com/labstack/echo/v4"

// Pipeline pipeline api operations
type Pipeline interface {
	Get(context echo.Context) error
	GetByProcessId(context echo.Context) error
	GetEvents(context echo.Context) error
}
