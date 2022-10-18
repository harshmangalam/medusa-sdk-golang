package orderedits

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type DeclineBody struct {
	DeclinedReason string `json:"declined_reason,omitempty"`
}

func NewDeclineBody() *DeclineBody {
	return new(DeclineBody)
}

func (d *DeclineBody) SetDeclineReason(reason string) *DeclineBody {
	d.DeclinedReason = reason
	return d
}

func (d *DeclineBody) Decline(id string, config *medusa.Config) (*OrderEdit, error) {
	path := fmt.Sprintf("/store/order-edits/%v/decline", id)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(d).Send(config)
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
