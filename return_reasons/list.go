package returnreasons

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

func List(config *medusa.Config) ([]*ReturnReason, error) {
	path := "/store/return-reasons"

	resp, err := request.
		NewRequest().
		SetMethod(http.MethodGet).
		SetPath(path).
		Send(config)

	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ReturnReasonListResponse)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.ReturnReasons, nil
}
