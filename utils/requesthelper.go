package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendGetRequest(url string, params url.Values) ([]byte, error) {
	url = fmt.Sprintf("%s?%s", url, params.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return checkResult(resp)
}

func checkResult(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != 200 {
		con, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("Request error:%s-%s", resp.Status, con))
	}
	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Result: %s. Error Msg: %s", string(result), err)
	}
	return result, nil
}

func PostRequest(url, payload string) (*http.Response, error) {
	body := &bytes.Buffer{}
	body.Write([]byte(payload))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}
