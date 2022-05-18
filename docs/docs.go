// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/applications": {
            "get": {
                "description": "Get all applications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "Get all applications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            },
            "post": {
                "description": "Update Application by company id and  repository id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "Update Application",
                "parameters": [
                    {
                        "description": "ApplicationWithUpdateOption Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "type": "string",
                        "description": "repository Id",
                        "name": "repositoryId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/applications/{id}": {
            "get": {
                "description": "Get application by appliction id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "Get application by appliction id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "application id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "repository id",
                        "name": "repositoryId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/bitbuckets/branches": {
            "get": {
                "description": "Gets Branches",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bitbucket"
                ],
                "summary": "Get Branches",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name",
                        "name": "userName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repository Id",
                        "name": "repoId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repository Name",
                        "name": "repoName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/bitbuckets/commits": {
            "get": {
                "description": "Gets Commit By Branch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bitbucket"
                ],
                "summary": "Get Commit By Branch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository Id",
                        "name": "repoId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch",
                        "name": "branch",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/companies": {
            "get": {
                "description": "Gets companies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get companies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Record count",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Loads Repositories",
                        "name": "loadRepositories",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Loads Applications",
                        "name": "loadApplications",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            },
            "post": {
                "description": "Saves company",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Save company",
                "parameters": [
                    {
                        "description": "Company data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/companies/{id}": {
            "get": {
                "description": "Gets company by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get company by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/companies/{id}/applications": {
            "get": {
                "description": "Gets RApplications by company id and repository type",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get Applications by company id and repository type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Record count",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/companies/{id}/repositories": {
            "get": {
                "description": "Gets RepositoriesDto by company id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get RepositoriesDto by company id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            },
            "put": {
                "description": "updates repositories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Update repositories by company id",
                "parameters": [
                    {
                        "description": "List Of Repositories data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Company id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Company Update Option",
                        "name": "companyUpdateOption",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/companies/{id}/repositories/{repoId}/webhooks": {
            "patch": {
                "description": "Update Webhook",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "Update Webhook to Enable or Disable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "action type [enable/disable]",
                        "name": "action",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repository type [github/bitbucket]",
                        "name": "repoType",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Company id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repository id",
                        "name": "repoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Webhook Id to disable webhook",
                        "name": "webhookId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/githubs/branches": {
            "get": {
                "description": "Gets Branches",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "Get Branches",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository Id",
                        "name": "repoId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/githubs/commits": {
            "get": {
                "description": "Gets commit by branch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "Get commit by branch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "branch",
                        "name": "branch",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repository Id",
                        "name": "repoId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/logs": {
            "post": {
                "description": "Stores logs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "Save log",
                "parameters": [
                    {
                        "description": "LogEvent Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/pipelines/ws": {
            "get": {
                "description": "Get events by process id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pipeline"
                ],
                "summary": "Get events by process id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "process id",
                        "name": "processId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/pipelines/{processId}": {
            "get": {
                "description": "Gets Pipeline or logs by pipeline processId [If action is \"get_pipeline\", then pipeline will be returned or logs will be returned. Available if local storage is enabled]",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pipeline"
                ],
                "summary": "Get Pipeline or logs [available if local storage is enabled]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pipeline ProcessId",
                        "name": "processId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "action",
                        "name": "action",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Record count",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/process_life_cycle_events": {
            "get": {
                "description": "Pulls auto trigger enabled steps",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessLifeCycle"
                ],
                "summary": "Pull Steps",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Agen name",
                        "name": "agent",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Pull size",
                        "name": "count",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Step type [BUILD, DEPLOY]",
                        "name": "step_type",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            },
            "post": {
                "description": "Stores process lifecycle event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessLifeCycle"
                ],
                "summary": "Save process lifecycle event",
                "parameters": [
                    {
                        "description": "ProcessLifeCycleEventList Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/processes": {
            "get": {
                "description": "Get Process List or count process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get Process List or count process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository Id",
                        "name": "repositoryId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "App Id",
                        "name": "appId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Commit Id",
                        "name": "commitId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Operation[countTodaysProcessByCompanyId]",
                        "name": "operation",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/processes/{processId}/steps/{step}/footmarks": {
            "get": {
                "description": "Get Footmarks By Process Id And Step",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get Footmarks By Process Id And Step",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Id",
                        "name": "processId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Step",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/processes_events": {
            "post": {
                "description": "Stores Pipeline process event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProcessEvent"
                ],
                "summary": "Save Pipeline process event",
                "parameters": [
                    {
                        "description": "PipelineProcessEvent Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/repositories/{id}": {
            "get": {
                "description": "Get repository by repository id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Repository"
                ],
                "summary": "Get repository by repository id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository id",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Load applications",
                        "name": "loadApplications",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/repositories/{id}/applications": {
            "get": {
                "description": "Get applications by repository id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Repository"
                ],
                "summary": "Get applications by repository id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository id",
                        "name": "repositoryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Record count",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.MetaData": {
            "type": "object",
            "properties": {
                "links": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                },
                "page": {
                    "type": "integer"
                },
                "page_count": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "common.ResponseDTO": {
            "type": "object",
            "properties": {
                "_metadata": {
                    "$ref": "#/definitions/common.MetaData"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "api-service API",
	Description:      "api-service  API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
