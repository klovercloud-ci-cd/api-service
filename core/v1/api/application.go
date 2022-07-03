package api

import (
	"github.com/labstack/echo/v4"
)

// Application application api operations
type Application interface {
	GetById(context echo.Context) error
	GetAll(context echo.Context) error
	CreatePipeline(context echo.Context) error
	UpdatePipeline(context echo.Context) error
}
