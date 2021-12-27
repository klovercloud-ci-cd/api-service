package main

import (
	"github.com/klovercloud-ci-cd/api-service/api"
	"github.com/klovercloud-ci-cd/api-service/config"
	_ "github.com/klovercloud-ci-cd/api-service/docs"
	"github.com/labstack/echo-contrib/jaegertracing"
)

// @title api-service API
// @description api-service  API
func main() {
	e := config.New()
	if config.EnableOpenTracing {
		c := jaegertracing.New(e, nil)
		defer c.Close()
	}
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//  goplantuml -recursive . > ClassDiagram.puml
// goreportcard-cli -v
// swag init --parseDependency --parseInternal

