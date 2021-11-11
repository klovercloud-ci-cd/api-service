package main

import (
	"github.com/klovercloud-ci-cd/api-service/api"
	"github.com/klovercloud-ci-cd/api-service/config"
	_ "github.com/klovercloud-ci-cd/api-service/docs"
)

// @title api-service API
// @description api-service  API
func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//  goplantuml -recursive . > ClassDiagram.puml
// goreportcard-cli -v
// swag init --parseDependency --parseInternal
