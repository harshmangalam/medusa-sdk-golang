package orderedits

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

func Retrieve(id string, config *medusa.Config) (*OrderEdit, error) {
	path := fmt.Sprintf("/store/order-edits/%v", id)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(RetrieveOrderEditResponse)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.OrderEdit, nil
}
