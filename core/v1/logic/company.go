package logic

import (
	"encoding/json"
	"fmt"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"log"
	"net/http"
)

type companyService struct {
	httpPublisher service.HttpClient
}

func (c companyService) GetApplicationsByRepositoryId(repoId string, companyId string, option v1.RepositoryQueryOption, status string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/repositories/"+repoId+"/applications?page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&companyId="+companyId+"&status="+status, header)

	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		return code, nil
	}
	return code, response
}

func (c companyService) GetApplicationsByCompanyIdAndRepositoryType(id string, _type string, option v1.CompanyQueryOption, status string) (httpCode int, data interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/companies/"+id+"/applications"+"?repository_type="+_type+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&status="+status, header)

	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		log.Println("[Error]", err.Error())
		return code, nil
	}
	return code, response
}

func (c companyService) GetApplicationByApplicationId(companyId string, repoId string, applicationId string) (httpCode int, data interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/applications/"+applicationId+"?companyId="+companyId+"&repositoryId="+repoId, header)

	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		log.Println("[Error]", err.Error())
		return code, nil
	}
	return code, response
}

func (c companyService) UpdateApplication(id string, repoId string, payload interface{}, option string) (httpCode int, err error) {
	marshal, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	post, err := c.httpPublisher.Post(config.KlovercloudIntegrationMangerUrl+"/applications?companyId="+id+"&repositoryId="+repoId+"&companyUpdateOption="+option, header, marshal)

	if err != nil {
		return 0, err
	}
	return post, nil
}

func (c companyService) UpdateRepositories(companyId string, company interface{}, option string) (httpCode int, err error) {
	marshal, err := json.Marshal(company)
	if err != nil {
		return 0, err
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	httpCode, err = c.httpPublisher.Put(config.KlovercloudIntegrationMangerUrl+"/companies/"+companyId+"/repositories?companyUpdateOption="+option, header, marshal)

	if err != nil {
		return 0, err
	}

	return httpCode, nil
}

func (c companyService) GetRepositoryByRepositoryId(id string, companyId string, loadApplications string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/repositories/"+id+"?loadApplications="+loadApplications+"&companyId="+companyId, header)

	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		return code, nil
	}
	return code, response
}

func (c companyService) Store(company interface{}) (httpCode int, error error) {
	marshal, marshalErr := json.Marshal(company)
	if marshalErr != nil {
		return http.StatusBadRequest, marshalErr
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, err := c.httpPublisher.Post(config.KlovercloudIntegrationMangerUrl+"/companies", header, marshal)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (c companyService) GetRepositoriesById(id string, option v1.CompanyQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/companies/"+id+"/repositories"+"?loadRepositories="+option.LoadRepositories+"&loadApplications="+option.LoadApplications, header)
	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		return code, nil
	}
	return code, response
}
func (c companyService) Get(option v1.CompanyQueryOption, status string) (httpCode int, body interface{}) {
	response := make(map[string]interface{})
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/companies"+"?loadRepositories="+option.LoadRepositories+"&loadApplications="+option.LoadApplications+"&status="+status, header)
	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		fmt.Println(er)
		return code, nil
	}
	return code, response
}
func (c companyService) GetById(headers map[string]string, id string, option v1.CompanyQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := c.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/companies/"+id+"?loadRepositories="+option.LoadRepositories+"&loadApplications="+option.LoadApplications, header)
	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		return code, nil
	}
	return code, response
}

// NewCompanyService returns Company type service
func NewCompanyService(publisher service.HttpClient) service.Company {
	return companyService{
		httpPublisher: publisher,
	}
}
