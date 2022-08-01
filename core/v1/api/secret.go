package api

import "github.com/labstack/echo/v4"

// Secret api operations
type Secret interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
