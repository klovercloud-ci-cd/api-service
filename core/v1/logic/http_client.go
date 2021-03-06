package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/opentracing"
	opentracer "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type httpClientService struct {
}

// Delete method that fires a Delete request.
func (h httpClientService) Delete(url string, header map[string]string) (httpCode int, err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
		log.Println("[DEBUG] Header: ", k, ":", v)
	}
	if err != nil {
		log.Println(err.Error())
	}
	client := &http.Client{}
	startTraceSpan(req, url, http.MethodDelete)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		var resBody common.ResponseDTO
		json.Unmarshal(body, &resBody)
		if err != nil {
			log.Println(err.Error())
			return resp.StatusCode, err
		} else {
			log.Println("[ERROR] Failed to communicate :", string(body))
			return resp.StatusCode, errors.New(resBody.Message)
		}
	}
	return resp.StatusCode, nil
}

// Put method that fires a Put request.
func (h httpClientService) Put(url string, header map[string]string, body []byte) (httpCode int, err error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	startTraceSpan(req, url, http.MethodPut)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Failed communicate :", err.Error())
		return resp.StatusCode, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		var resBody common.ResponseDTO
		json.Unmarshal(body, &resBody)
		if err != nil {
			log.Println("[ERROR] Failed communicate ", err.Error())
			return resp.StatusCode, err
		} else {
			log.Println("[SUCCESS] Successful :", string(body))
			return resp.StatusCode, errors.New(resBody.Message)
		}
	}
	return resp.StatusCode, nil
}

// Get method that fires a get request.
func (h httpClientService) Get(url string, header map[string]string) (httpCode int, body []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil, err
	}
	startTraceSpan(req, url, http.MethodGet)

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

// Post method that fires a Post request.
func (h httpClientService) Post(url string, header map[string]string, body []byte) (httpCode int, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	startTraceSpan(req, url, http.MethodPost)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Failed communicate :", err.Error())
		return http.StatusBadRequest, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		var resBody common.ResponseDTO
		json.Unmarshal(body, &resBody)
		if err != nil {
			log.Println("[ERROR] Failed to communicate ", err.Error())
			return resp.StatusCode, err
		} else {
			log.Println("[ERROR] Failed to communicate :", string(body))
			return resp.StatusCode, errors.New(resBody.Message)
		}
	}
	return resp.StatusCode, nil
}

// Patch method that fires a Patch request.
func (h httpClientService) Patch(url string, header map[string]string, body []byte) (httpCode int, err error) {
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	for k, v := range header {
		req.Header.Set(k, v)
		log.Println("[DEBUG] Header: ", k, ":", v)
	}
	if err != nil {
		log.Println(err.Error())
	}
	client := &http.Client{}
	startTraceSpan(req, url, "PATCH")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return resp.StatusCode, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		var resBody common.ResponseDTO
		json.Unmarshal(body, &resBody)
		if err != nil {
			log.Println(err.Error())
			return resp.StatusCode, err
		} else {
			log.Println("[ERROR] Failed to communicate :", string(body))
			return resp.StatusCode, errors.New(resBody.Message)
		}
	}
	return resp.StatusCode, nil
}

// startTraceSpan starts a span
func startTraceSpan(req *http.Request, url, httpMethod string) {
	if config.EnableOpenTracing {
		span, _ := opentracer.StartSpanFromContext(context.Background(), "client")
		ext.SpanKindRPCClient.Set(span)
		ext.HTTPUrl.Set(span, url)
		ext.HTTPMethod.Set(span, httpMethod)
		defer span.Finish()
		if err := opentracing.Inject(span, req); err != nil {
			log.Println(err.Error())
		}
	}
}

// NewHttpClientService returns HttpClient type service
func NewHttpClientService() service.HttpClient {
	return &httpClientService{}
}
