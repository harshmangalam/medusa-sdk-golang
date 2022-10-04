package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
)

type Request struct {
	Method string `json:"method"`
	Data   any    `json:"data"`
	Path   string `json:"params"`
}

func NewRequest() *Request {
	return new(Request)
}

func (req *Request) SetMethod(method string) *Request {
	req.Method = method
	return req
}
func (req *Request) SetData(data any) *Request {
	req.Data = data
	return req
}

func (req *Request) SetPath(path string) *Request {
	req.Path = path
	return req
}

func (req *Request) Send(medusa *medusa.Medusa) (*http.Response, error) {
	url := medusa.BaseUrl + req.Path
	client := &http.Client{}
	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}

	switch req.Method {
	case http.MethodGet:
		httpReq, err := http.NewRequest(http.MethodGet, url, nil)

		fmt.Println(httpReq.Cookies())
		if err != nil {
			return nil, err
		}
		httpReq.Header = headers

		httpReq.AddCookie(medusa.Cookie)

		resp, err := client.Do(httpReq)

		if err != nil {
			return nil, err
		}

		return resp, nil

	case http.MethodPost:
		jsonData, err := json.Marshal(req.Data)
		if err != nil {
			return nil, err
		}

		buff := bytes.NewBuffer(jsonData)
		httpReq, err := http.NewRequest(http.MethodPost, url, buff)

		if err != nil {
			return nil, err
		}

		httpReq.Header = headers
		resp, err := client.Do(httpReq)

		if err != nil {
			return nil, err
		}

		return resp, nil

	default:
		err := errors.New("request method is invalid")
		return nil, err
	}

}
