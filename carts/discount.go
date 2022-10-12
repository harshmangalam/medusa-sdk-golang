package carts

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

func DeleteDiscount(cartId string, code string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/discounts/%v", cartId, code)
	resp, err := request.NewRequest().SetMethod(http.MethodDelete).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
