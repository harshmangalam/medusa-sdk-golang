package auth

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/customers"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type GetSessionData struct {
	Customer customers.Customer `json:"customer,omitempty"`
}

type GetSessionResponse struct {
	Data  *GetSessionData
	Error *response.Error
}

func GetSession(config *medusa.Config) (*GetSessionResponse, error) {
	path := "/store/auth/"
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(GetSessionResponse)

	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(GetSessionData)
		if json.Unmarshal(body, respData); err != nil {
			return nil, err
		}
		respBody.Data = respData

	case http.StatusUnauthorized:
		respErr := new(response.Error)
		respErr.Message = "Unauthorized"
		respBody.Error = respErr

	default:
		respErr := new(response.Error)
		if json.Unmarshal(body, respErr); err != nil {
			return nil, err
		}
		respBody.Error = respErr

	}

	return respBody, err
}
