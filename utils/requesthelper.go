package utils

import (
	"fmt"
	"net/http"
	"net/url"
)

func sendGetRequest(url string, params url.Values) ([]byte, error) {
	url := fmt.Sprintf("%s?%s", url, params.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}
