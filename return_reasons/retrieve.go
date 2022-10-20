package returnreasons

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

func Retrieve(id string, config *medusa.Config) (*ReturnReason, error) {
	path := fmt.Sprintf("/store/return-reasons/%v", id)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(ReturnReasonResponse)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.ReturnReason, nil
}
