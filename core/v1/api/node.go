package api

import "github.com/labstack/echo/v4"

// Node api operations
type Node interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
