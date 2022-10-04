package utils

import (
	"io/ioutil"
	"net/http"
)

func ParseResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
