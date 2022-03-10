package api

import "github.com/labstack/echo/v4"

// Process process api operations
type Process interface {
	Get(context echo.Context) error
}
