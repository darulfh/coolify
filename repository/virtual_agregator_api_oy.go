package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/model"
)

type VirtualAgregatorOyApi interface {
	GenerateVaApi(payload model.GenerateVirtualAgregator) (*model.VaNumber, error)
	GetVaIdStatusVaApi(virtualId string) (*model.VaNumber, error)
}

type virtualAgregatorOyApiRepository struct{}

func NewVirtualAgregatorOyApiRepository() VirtualAgregatorOyApi {
	return &virtualAgregatorOyApiRepository{}
}

func (*virtualAgregatorOyApiRepository) GenerateVaApi(payload model.GenerateVirtualAgregator) (*model.VaNumber, error) {
	url := fmt.Sprintf("%s/generate-static-va", config.AppConfig.BaseUrl)
	resp, err := doRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var apiResponse model.VaNumber
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &apiResponse, nil
}

func (*virtualAgregatorOyApiRepository) GetVaIdStatusVaApi(virtualId string) (*model.VaNumber, error) {
	url := fmt.Sprintf("%s/static-virtual-account%s", config.AppConfig.BaseUrl, virtualId)
	resp, err := doRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var vaResponse model.VaNumber
	if err := json.Unmarshal(body, &vaResponse); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &vaResponse, nil
}
