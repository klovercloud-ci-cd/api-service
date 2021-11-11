package logic

import (
	"bytes"
	"errors"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type httpClientService struct {
}

func (h httpClientService) Put(url string, header map[string]string, body []byte) (httpCode int, err error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Failed communicate :", err.Error())
		return http.StatusBadRequest, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, err
			log.Println("[ERROR] Failed communicate ", err.Error())
		} else {
			log.Println("[SUCCESS] Successful :", string(body))
		}
	}
	return resp.StatusCode, nil
}

func (h httpClientService) Get(url string, header map[string]string) (httpCode int, body []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return res.StatusCode, nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err.Error())
			return res.StatusCode, nil, err
		}
		return res.StatusCode, jsonDataFromHttp, nil
	}
	return res.StatusCode, nil, errors.New("Status: " + res.Status + ", code: " + strconv.Itoa(res.StatusCode))
}

func (h httpClientService) Post(url string, header map[string]string, body []byte) (httpCode int, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Failed communicate :", err.Error())
		return http.StatusBadRequest, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, err
			log.Println("[ERROR] Failed communicate ", err.Error())
		} else {
			log.Println("[ERROR] Failed communicate :", string(body))
		}
	}
	return resp.StatusCode, nil
}

// NewHttpClientService returns HttpClient type service
func NewHttpClientService() service.HttpClient {
	return &httpClientService{}
}
