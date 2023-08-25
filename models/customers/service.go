package customers

import (
	"context"
	"encoding/json"
	"net/http"
	"service-mini-restapi/helper"
)

type service struct {
	repository Repository
}

type Service interface {
	GetFromApi(ctx context.Context) ([]CustomerApi, error)
	Create(input RequestCustomer) (ResultRequest, error)
	GetAll(search SearchCustomer) ([]Customers, error)
	Delete(input RequestDelete) error
	Update(input RequestUpdate) (ResultUpdate, error)
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

type ResponseBodyData struct {
	ResponseCode  string      `json:"responseCode"`
	ResponseDesc  string      `json:"responseDesc"`
	ResponseData  []CustomerApi `json:"responseData"`
	ResponseError string      `json:"responseError"`
}

func (s *service) GetFromApi(ctx context.Context) ([]CustomerApi, error) {
	var err error
	var results []CustomerApi

	var responseBody ResponseBodyData

	// Make the HTTP GET request with the desired header
	req, err := http.NewRequest("GET", "https://mocki.io/v1/34f50297-f497-498a-888e-0f24e7eb97ec", nil)
	if err != nil {
		return results, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Security-Policy", "default-src 'self'")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return results, err
	}
	return responseBody.ResponseData, nil
}

func (s *service) Create(input RequestCustomer) (ResultRequest, error) {
	var err error
	var result ResultRequest
	result, err = s.repository.Create(input)
	helper.PanicIfError(err)
	return result, nil
}

func (s *service) GetAll(search SearchCustomer) ([]Customers, error) {
	var err error
	var results []Customers
	results, err = s.repository.GetAll(search)
	helper.PanicIfError(err)
	return results, nil
}

func (s *service) Delete(input RequestDelete) error {
	var err error
	err = s.repository.Delete(input)
	helper.PanicIfError(err)
	return nil
}

func (s *service) Update(input RequestUpdate) (ResultUpdate, error) {
	var err error
	var result ResultUpdate

	result, err = s.repository.Update(input)
	helper.PanicIfError(err)
	return result, nil

}
