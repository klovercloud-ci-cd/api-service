package v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
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

var (
	upgrader = websocket.Upgrader{}
)

func (p pipelineApi) GetEvents(context echo.Context) error {
	processId := context.QueryParam("processId")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {

		}
	}(ws)

	status := make(chan map[string]interface{})
	for {
		go p.pipelineService.ReadEventsByProcessId(status, processId)
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

// Get... Get logs [available if local storage is enabled]
// @Summary Get Logs [available if local storage is enabled]
// @Description Gets logs by pipeline processId [available if local storage is enabled]
// @Tags Pipeline
// @Produce json
// @Param processId path string true "Pipeline ProcessId"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/pipelines/{processId} [GET]
func (p pipelineApi) GetLogs(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
		if err := checkAuthority(userResourcePermission, string(enums.PIPELINE), "", string(enums.READ)); err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
	}
	option := getPipelineQueryOption(context)
	code, data := p.pipelineService.GetByProcessId(id, option)
	return context.JSON(code, data)
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
