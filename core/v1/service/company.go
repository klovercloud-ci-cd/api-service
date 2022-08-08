package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Company Company operations.
type Company interface {
	GetById(headers map[string]string, id, action string, option v1.CompanyQueryOption) (httpCode int, body interface{})
	Get(option v1.CompanyQueryOption, status string) (httpCode int, body interface{})
	GetRepositoriesById(id string, option v1.CompanyQueryOption) (httpCode int, body interface{})
	Store(company interface{}) (httpCode int, error error)
	GetRepositoryByRepositoryId(id string, companyId string, loadApplications string) (httpCode int, body interface{})
	GetApplicationsByRepositoryId(repoId string, companyId string, option v1.RepositoryQueryOption, status string) (httpCode int, body interface{})
	UpdateRepositories(companyId string, company interface{}, option string) (httpCode int, error error)
	UpdateApplications(id string, repoId string, payload interface{}, option, validity string) (httpCode int, error error)
	GetApplicationByApplicationId(companyId string, repoId string, applicationId string) (httpCode int, data interface{})
	GetAllApplications(companyId string, option v1.CompanyQueryOption) (httpCode int, data interface{})
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type string, option v1.CompanyQueryOption, status string) (httpCode int, data interface{})
	UpdateWebhook(id, repoId, appId, url, webhookId, action, repoType string) (httpCode int, error error)
	CreateApplicationPipeline(companyId, repositoryId, appId string, payload interface{}) (httpCode int, body interface{})
	UpdateApplicationPipeline(companyId, repositoryId, appId string, payload interface{}) (httpCode int, body interface{})
}
