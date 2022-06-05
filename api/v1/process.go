package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type processApi struct {
	processService service.Process
	jwtService     service.Jwt
}

// Get... Get Process by Id
// @Summary Get Process by Id
// @Description Get Process by Id
// @Tags Process
// @Produce json
// @Param processId path string true "ProcessId"
// @Success 200 {object} common.ResponseDTO{v1.Process}
// @Router /api/v1/processes/{processId} [GET]
func (p processApi) GetById(context echo.Context) error {
	processId := context.Param("processId")
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	code, data := p.processService.GetById(companyId, processId)
	if code != 200 {
		return common.GenerateErrorResponse(context, "[ERROR]: Process not found", "Operation failed")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Operation Succesful")
}

// GetLogsByProcessIdAndStepAndFootmark... Get logs by Footmarks, Process Id And Step
// @Summary Get logs by Footmarks, Process Id And Step
// @Description Get logs by Footmarks, Process Id And Step
// @Tags Process
// @Produce json
// @Param processId path string true "Process Id"
// @Param step path string true "Step"
// @Param footmark path string true "Footmark"
// @Param claims query string true "Claims"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/processes/{processId}/steps/{step}/footmarks/{footmark}/logs [GET]
func (p processApi) GetLogsByProcessIdAndStepAndFootmark(context echo.Context) error {
	processId := context.Param("processId")
	step := context.Param("step")
	footmark := context.Param("footmark")
	claims := context.QueryParam("claims")
	option := getQueryOption(context)
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	code, res := p.processService.GetLogsByProcessIdAndStepAndFootmark(processId, step, footmark, claims, option)
	if code != 200 {
		return common.GenerateErrorResponse(context, "", "Logs not found")
	}
	return context.JSON(code, res)
}

// GetFootmarksByProcessIdAndStep... Get Footmarks By Process Id And Step
// @Summary Get Footmarks By Process Id And Step
// @Description Get Footmarks By Process Id And Step
// @Tags Process
// @Produce json
// @Param processId path string true "Process Id"
// @Param step path string true "Step"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/processes/{processId}/steps/{step}/footmarks [GET]
func (p processApi) GetFootmarksByProcessIdAndStep(context echo.Context) error {
	process := context.Param("processId")
	step := context.Param("step")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	code, res := p.processService.GetFootmarksByProcessIdAndStep(process, step)
	if code != 200 {
		return common.GenerateErrorResponse(context, "", "Footmarks not found")
	}
	return context.JSON(code, res)
}

// Get... Get Process List or count process
// @Summary Get Process List or count process
// @Description Get Process List or count process
// @Tags Process
// @Produce json
// @Param repositoryId query string false "Repository Id"
// @Param appId query string false "App Id"
// @Param commitId query string false "Commit Id"
// @Param operation query string false "Operation[countTodaysProcessByCompanyId]"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/processes [GET]
func (p processApi) Get(context echo.Context) error {
	var companyId string
	repositoryId := context.QueryParam("repositoryId")
	appId := context.QueryParam("appId")
	commitId := context.QueryParam("commitId")
	option := getProcessQueryOption(context)
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	code, data := p.processService.Get(companyId, repositoryId, appId, commitId, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Processes Query Failed", "Operation Failed")
}

//this function is for set all query param
func getProcessQueryOption(context echo.Context) v1.ProcessQueryOption {
	option := v1.ProcessQueryOption{}
	option.Pagination.Page = context.QueryParam("page")
	option.Pagination.Limit = context.QueryParam("limit")
	option.Step = context.QueryParam("step")
	return option
}

// NewProcessApi returns Process type api
func NewProcessApi(processService service.Process, jwtService service.Jwt) api.Process {
	return &processApi{
		processService: processService,
		jwtService:     jwtService,
	}
}
