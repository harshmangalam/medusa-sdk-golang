package carts

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

func Complete(cart_id string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/complete", cart_id)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
