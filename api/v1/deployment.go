package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type deploymentApi struct {
	deploymentService service.Deployment
	jwtService        service.Jwt
	processService    service.Process
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all Deployment by agent name, owner reference and process
// @Tags Deployment
// @Produce json
// @Param owner-reference path string false "Owner Reference"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/deployments [GET]
func (d deploymentApi) Get(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, d.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not given", "Operation Failed")
	}
	agent := context.QueryParam("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not given", "Operation Failed")
	}
	ownerReference := context.QueryParam("owner-reference")
	code, jsonBody := d.processService.GetById(companyId, processId)
	if code != 200 {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	body, err := json.Marshal(jsonBody)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	responseDTO := common.ResponseDTO{}
	if err = json.Unmarshal(body, &responseDTO); err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	var process v1.Process
	processBody, err := json.Marshal(responseDTO.Data)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	if err = json.Unmarshal(processBody, &process); err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	if process.ProcessId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process not found", "Operation Failed")
	}
	option := getK8sObjectQueryOption(context)
	code, data := d.deploymentService.Get(agent, ownerReference, processId, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
}

// Get... Get by ID Api
// @Summary Get by ID api
// @Description Api for getting a Deployments by id, agent name, and process id
// @Tags Deployment
// @Produce json
// @Param id query string true "ID"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Success 200 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/deployments/{id} [GET]
func (d deploymentApi) GetByID(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, d.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not given", "Operation Failed")
	}
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: ID is not given", "Operation Failed")
	}
	agent := context.QueryParam("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not given", "Operation Failed")
	}
	code, jsonBody := d.processService.GetById(companyId, processId)
	if code != 200 {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	body, err := json.Marshal(jsonBody)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	responseDTO := common.ResponseDTO{}
	if err = json.Unmarshal(body, &responseDTO); err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	var process v1.Process
	processBody, err := json.Marshal(responseDTO.Data)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	if err = json.Unmarshal(processBody, &process); err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
	}
	if process.ProcessId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process not found", "Operation Failed")
	}
	code, data := d.deploymentService.GetByID(id, agent, processId)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "[ERROR]: k8s Object Query Failed", "Operation Failed")
}

// NewDeploymentApi returns api.Deployment type api
func NewDeploymentApi(deploymentService service.Deployment, jwtService service.Jwt, processService service.Process) api.Deployment {
	return &deploymentApi{
		deploymentService: deploymentService,
		jwtService:        jwtService,
		processService:    processService,
	}
}
