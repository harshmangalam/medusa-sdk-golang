package swaps

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveSwapData struct {
	Swap *schema.Swap `json:"swap"`
}

type RetrieveSwapResponse struct {
	// Success response
	Data *RetrieveSwapData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func RetrieveByCartId(cartId string, config *medusa.Config) (*RetrieveSwapResponse, error) {
	path := fmt.Sprintf("/store/swaps/%v", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

}
