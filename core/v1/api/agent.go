package api

import "github.com/labstack/echo/v4"

// Agent operations
type Agent interface {
	Save(context echo.Context) error
	Get(context echo.Context) error
}
