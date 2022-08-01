package api

import "github.com/labstack/echo/v4"

// Certificate api operations
type Certificate interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
