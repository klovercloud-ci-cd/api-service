package api

import "github.com/labstack/echo/v4"

// Company company api operations
type Company interface {
	Save(context echo.Context) error
	GetById(context echo.Context) error
	Get(context echo.Context) error
	GetRepositoriesById(context echo.Context) error
	UpdateRepositories(context echo.Context) error
	UpdateApplications(context echo.Context) error
	GetApplicationsByCompanyIdAndRepositoryType(context echo.Context) error
	UpdateWebhook(context echo.Context) error
}
