package repository

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/model"
	"github.com/google/uuid"
)

type IakApiRepository interface {
	IakTopUpPayRepository(payload *model.PrePaidIakBody) (*model.PrePaidIakResponse, error)

	BpjsInquiryRepository(payload *model.IakInquiryBody) (*model.BpjsIAKResponse, error)
	BpjsPayRepository(payload *model.IakPayBody) (*model.BpjsIAKResponse, error)
	BpjsCheckRepository(payload *model.IakPayBody) (*model.BpjsIAKResponse, error)

	ElectricityBillInquiryRepository(payload *model.IakInquiryBody) (*model.IakPostPaidResponse, error)
	ElectricityBillPayRepository(payload *model.IakPayBody) (*model.IakPostPaidResponse, error)
	ElectricityBillCheckRepository(payload *model.IakPayBody) (*model.IakPostPaidResponse, error)

	ElectricityTokenInquiryRepository(payload *model.PrePaidIakBody) (*model.IakElectricityTokenInquiry, error)
}

type iakApiRepository struct{}

func NewIakApiRepository() IakApiRepository {
	return &iakApiRepository{}
}

func (*iakApiRepository) IakTopUpPayRepository(payload *model.PrePaidIakBody) (*model.PrePaidIakResponse, error) {

	payload.Sign = sign(payload.RefID)

	// resp, err := doRequestIak(http.MethodPost, "https://prepaid.iak.dev/api/top-up", payload)
	resp, err := doRequestIak(http.MethodPost, "https://prepaid.iak.id/api/top-up", payload)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.PrePaidIakResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &response, nil
}

func (*iakApiRepository) BpjsInquiryRepository(payload *model.IakInquiryBody) (*model.BpjsIAKResponse, error) {
	payload.Commands = "inq-pasca"
	payload.Code = "BPJS"
	payload.RefID = uuid.New().String()
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign(payload.RefID)

	resp, err := doRequestIak(http.MethodPost, config.AppConfig.BaseUrlIakPostPaid+"/api/v1/bill/check", payload)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.BpjsIAKResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	if response.Data.Message != "INQUIRY SUCCESS" {
		return nil, fmt.Errorf(response.Data.Message)
	}

	return &response, nil
}

func (*iakApiRepository) BpjsPayRepository(payload *model.IakPayBody) (*model.BpjsIAKResponse, error) {
	payload.Commands = "pay-pasca"
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign(strconv.Itoa(payload.TrID))

	resp, err := doRequestIak(http.MethodPost, config.AppConfig.BaseUrlIakPostPaid+"/api/v1/bill/check", payload)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.BpjsIAKResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	if response.Data.Message != "PAYMENT SUCCESS" {
		return nil, fmt.Errorf(response.Data.Message)
	}

	return &response, nil
}

func (*iakApiRepository) BpjsCheckRepository(payload *model.IakPayBody) (*model.BpjsIAKResponse, error) {
	payload.Commands = "checkstatus"
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign("cs")

	fmt.Printf("response123123: %+v\n", payload)

	resp, err := doRequestIak(http.MethodPost, config.AppConfig.BaseUrlIakPostPaid+"/api/v1/bill/check", payload)

	fmt.Printf("response0: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response model.BpjsIAKResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response1: %+v\n", response)

	return &response, nil
}

func (*iakApiRepository) ElectricityBillInquiryRepository(payload *model.IakInquiryBody) (*model.IakPostPaidResponse, error) {
	payload.Commands = "inq-pasca"
	payload.Code = "PLNPOSTPAID"
	payload.RefID = uuid.New().String()
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign(payload.RefID)

	resp, err := doRequestIak(http.MethodPost, config.AppConfig.BaseUrlIakPostPaid+"/api/v1/bill/check", payload)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.IakPostPaidResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	if response.Data.Message != "INQUIRY SUCCESS" {
		return nil, fmt.Errorf(response.Data.Message)
	}

	return &response, nil
}
func (*iakApiRepository) ElectricityBillPayRepository(payload *model.IakPayBody) (*model.IakPostPaidResponse, error) {
	payload.Commands = "pay-pasca"
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign(strconv.Itoa(payload.TrID))

	resp, err := doRequestIak(http.MethodPost, config.AppConfig.BaseUrlIakPostPaid+"/api/v1/bill/check", payload)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.IakPostPaidResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	if response.Data.Message != "PAYMENT SUCCESS" {
		return nil, fmt.Errorf(response.Data.Message)
	}

	return &response, nil
}
func (*iakApiRepository) ElectricityBillCheckRepository(payload *model.IakPayBody) (*model.IakPostPaidResponse, error) {
	payload.Commands = "checkstatus"
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign("cs")

	fmt.Printf("response123123: %+v\n", payload)

	resp, err := doRequestIak(http.MethodPost, config.AppConfig.BaseUrlIakPostPaid+"/api/v1/bill/check", payload)

	fmt.Printf("response0: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response model.IakPostPaidResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response1: %+v\n", response)

	return &response, nil
}

func (*iakApiRepository) ElectricityTokenInquiryRepository(payload *model.PrePaidIakBody) (*model.IakElectricityTokenInquiry, error) {
	payload.Username = config.AppConfig.UsernameIak
	payload.Sign = sign(payload.CustomerID)

	// resp, err := doRequestIak(http.MethodPost, "https://prepaid.iak.dev/api/inquiry-pln", payload)
	resp, err := doRequestIak(http.MethodPost, "https://prepaid.iak.id/api/inquiry-pln", payload)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.IakElectricityTokenInquiry
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	if response.Data.Message != "SUCCESS" {
		return nil, fmt.Errorf(response.Data.Message)

	}

	return &response, nil

}

func sign(id string) string {

	apiKey := config.AppConfig.ApiKeyIak
	username := config.AppConfig.UsernameIak

	hash := md5.New()
	_, _ = hash.Write([]byte(username + apiKey + id))

	md5 := hash.Sum(nil)

	return fmt.Sprintf("%x", md5)
}
