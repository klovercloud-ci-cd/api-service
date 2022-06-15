package api

import "github.com/labstack/echo/v4"

type ProcessEvent interface {
	Get(context echo.Context) error
	Save(context echo.Context) error
}
