package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Company Company operations.
type Company interface {
	GetCompaniesById(headers map[string]string, id string, option v1.CompanyQueryOption) (httpCode int, body interface{})
	GetCompanies(option v1.CompanyQueryOption, status string) (httpCode int, body interface{})
	GetRepositoriesById(id string, option v1.CompanyQueryOption) (httpCode int, body interface{})
	Store(company interface{}) (httpCode int, error error)
	GetRepositoryByRepositoryId(id string, companyId string, loadApplications string) (httpCode int, body interface{})
	GetApplicationsByRepositoryId(repoId string, companyId string, option v1.RepositoryQueryOption, status string) (httpCode int, body interface{})
	UpdateRepositories(companyId string, company interface{}, option string) (httpCode int, error error)
	UpdateApplication(id string, repoId string, payload interface{}, option string) (httpCode int, error error)
	GetApplicationByApplicationId(companyId string, repoId string, applicationId string) (httpCode int, data interface{})
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type string, option v1.CompanyQueryOption, status string) (httpCode int, data interface{})
}
