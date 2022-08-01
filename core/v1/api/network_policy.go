package api

import "github.com/labstack/echo/v4"

// NetworkPolicy api operations
type NetworkPolicy interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
