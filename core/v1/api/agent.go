package api

import "github.com/labstack/echo/v4"

// Agent operations
type Agent interface {
	Get(context echo.Context) error
	GetTerminalByName(context echo.Context) error
	Save(context echo.Context) error
}
