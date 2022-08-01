package api

import "github.com/labstack/echo/v4"

// StatefulSet api operations
type StatefulSet interface {
	Get(context echo.Context) error
	GetByID(context echo.Context) error
}
