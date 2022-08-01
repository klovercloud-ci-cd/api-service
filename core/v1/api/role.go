package api

import "github.com/labstack/echo/v4"

// Role api operations
type Role interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
