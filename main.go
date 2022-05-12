package main

import (
	"github.com/klovercloud-ci-cd/api-service/api"
	"github.com/klovercloud-ci-cd/api-service/config"
	_ "github.com/klovercloud-ci-cd/api-service/docs"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// @title api-service API
// @description api-service  API
func main() {
	e := config.New()
	if config.EnableOpenTracing {
		c := jaegertracing.New(e, nil)
		defer c.Close()
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
