package api

import "github.com/labstack/echo/v4"

type LogEvent interface {
	Save(context echo.Context) error
}
