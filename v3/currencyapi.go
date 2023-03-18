package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const BaseUrl = "https://api.currencyapi.com/v3"

type httpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

func call(key string, endpoint string, client httpClient, dest any) error {
	var req, err1 = http.NewRequest("GET", endpoint, nil)
	if err1 != nil {
		return fmt.Errorf("currencyapi: creation of the http request failed with an error: %+v", err1)
	}

	req.Header.Set("apikey", key)

	var response, err2 = client.Do(req)
	if err2 != nil {
		return fmt.Errorf("currencyapi: http request failed with an error: %+v", err2)
	}

	defer response.Body.Close()

	var body, err3 = io.ReadAll(response.Body)
	if err3 != nil {
		return fmt.Errorf("currencyapi: reading response body failed with an error: %+v", err3)
	}

	var err4 = json.Unmarshal(body, dest)
	if err4 != nil {
		return fmt.Errorf("currencyapi: unmarshal json failed with an error: %+v", err4)
	}

	return nil
}

type ClientV3 struct {
	httpClient httpClient
	key        string
	baseURL    *url.URL
}

func NewClient(key string, httpClient httpClient) (*ClientV3, error) {
	if len(key) == 0 {
		return nil, errors.New("key cannot be empty")
	}
	if httpClient == nil {
		return nil, errors.New("http client cannot be nil")
	}

	var u, err0 = url.Parse(BaseUrl)
	if err0 != nil {
		panic(err0)
	}

	var client = new(ClientV3)
	client.httpClient = httpClient
	client.key = key
	client.baseURL = u
	return client, nil
}

func (c *ClientV3) Status() (*StatusResponse, error) {
	var response = new(StatusResponse)
	var u = c.baseURL.JoinPath("/status")
	return response, call(c.key, u.String(), c.httpClient, response)
}

func (c *ClientV3) Latest(request LatestRequest) (*LatestResponse, error) {
	var response = new(LatestResponse)
	var u = c.baseURL.JoinPath("/latest")
	var q = u.Query()
	q.Add("base_currency", request.From)
	q.Add("currencies", strings.Join(request.To, ","))
	u.RawQuery = q.Encode()
	return response, call(c.key, u.String(), c.httpClient, response)
}

//func Currencies(params map[string]string) []byte {
//	return apiCall("currencies", params)
//}
//
//func Historical(params map[string]string) []byte {
//	return apiCall("historical", params)
//}
//
//func Range(params map[string]string) []byte {
//	return apiCall("range", params)
//}
//
//func Convert(params map[string]string) []byte {
//	return apiCall("convert", params)
//}
