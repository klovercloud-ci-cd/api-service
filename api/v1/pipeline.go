package v1

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type pipelineApi struct {
	pipelineService service.Pipeline
	jwtService      service.Jwt
}

// Create... Create pipeline
// @Summary  Create Pipeline
// @Description Create Pipeline by repository id, application url
// @Tags Pipeline
// @Accept json
// @Produce json
// @Param pipeline body interface{} true "pipeline"
// @Param repositoryId query string true "Repository id"
// @Param url query string true "Url"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/pipelines [POST]
func (p pipelineApi) Create(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.CREATE)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	repoId := context.QueryParam("repositoryId")
	url := context.QueryParam("url")
	code, res := p.pipelineService.Create(companyId, repoId, url, formData)
	if code == 200 {
		return context.JSON(http.StatusOK, res)
	}
	return common.GenerateErrorResponse(context, "pipeline creation failed", "operation failed")
}

// Update... Update pipeline
// @Summary  Update Pipeline
// @Description Update Pipeline by repository id, application url
// @Tags Pipeline
// @Accept json
// @Produce json
// @Param pipeline body interface{} true "pipeline"
// @Param repositoryId query string true "Repository id"
// @Param url query string true "Url"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/pipelines [PUT]
func (p pipelineApi) Update(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.UPDATE)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	repoId := context.QueryParam("repositoryId")
	url := context.QueryParam("url")
	code, res := p.pipelineService.Update(companyId, repoId, url, formData)
	if code == 200 {
		return context.JSON(http.StatusOK, res)
	}
	return common.GenerateErrorResponse(context, "pipeline update failed", "operation failed")
}

var (
	upgrader = websocket.Upgrader{}
)

// Get... Get pipeline for validation
// @Summary  Get Pipeline for validation
// @Description Get Pipeline for validation by repository id, application url and revision
// @Tags Pipeline
// @Accept json
// @Produce json
// @Param action query string true "action [GET_PIPELINE_FOR_VALIDATION/dashboard_data"]"
// @Param repositoryId query string true "repository id"
// @Param url query string true "application url"
// @Param revision query string true "commit id or branch name"
// @Param from query string false "From Data"
// @Param to query string false "To Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/pipelines [GET]
func (p pipelineApi) Get(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	action := context.QueryParam("action")
	repoId := context.QueryParam("repositoryId")
	url := context.QueryParam("url")
	revision := context.QueryParam("revision")
	from := context.QueryParam("from")
	to := context.QueryParam("to")
	code, data := p.pipelineService.Get(companyId, repoId, url, revision, action, from, to)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Pipeline Query Failed", "Operation Failed")
}

// Get.. Get events by process id
// @Summary Get events by process id
// @Description Get events by process id
// @Tags Pipeline
// @Produce json
// @Param processId query string true "company_id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/pipelines/ws [GET]
func (p pipelineApi) GetEvents(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	userId := context.QueryParam("userId")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
		userId = userResourcePermission.UserId
	}

	status := make(chan map[string]interface{})
	for {
		go p.pipelineService.ReadEventsByCompanyId(status, companyId, userId)
		jsonStr, err := json.Marshal(<-status)
		if err != nil {
			log.Println(err.Error())
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte(jsonStr))
		if err != nil {
			log.Println("[ERROR]: Failed to write", err.Error())
			err := ws.Close()
			if err != nil {
				return err
			}
		}
		_, _, err = ws.ReadMessage()
		if err != nil {
			log.Println("[ERROR]: Failed to read", err.Error())
			err := ws.Close()
			if err != nil {
				return err
			}
		}

	}
}

// Get... Get Pipeline or logs [available if local storage is enabled]
// @Summary Get Pipeline or logs [available if local storage is enabled]
// @Description Gets Pipeline or logs by pipeline processId [If action is "get_pipeline", then pipeline will be returned or logs will be returned. Available if local storage is enabled]
// @Tags Pipeline
// @Produce json
// @Param id path string true "Pipeline ProcessId"
// @Param action query int64 false "action"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/pipelines/{processId} [GET]
func (p pipelineApi) GetByProcessId(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no processId is provided", "Please provide processId")
	}
	action := context.QueryParam("action")
	companyId := context.QueryParam("companyId")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	option := getPipelineQueryOption(context)
	code, data := p.pipelineService.GetByProcessId(companyId, id, action, option)
	if code == 200 {
		return context.JSON(http.StatusOK, data)
	}
	return common.GenerateErrorResponse(context, data, "Operation Failed")
}

func getPipelineQueryOption(context echo.Context) v1.Pagination {
	option := v1.Pagination{}
	option.Page = context.QueryParam("page")
	option.Limit = context.QueryParam("limit")
	return option
}

// NewPipelineApi returns Pipeline type api
func NewPipelineApi(pipelineService service.Pipeline, jwtService service.Jwt) api.Pipeline {
	return &pipelineApi{
		pipelineService: pipelineService,
		jwtService:      jwtService,
	}
}
