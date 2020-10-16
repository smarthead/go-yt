package rest

import (
	"bytes"
	"encoding/json"
	"go-yt/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

type Client struct {
	YoutrackUrl *url.URL
	Token       string

	DefaultHeaders map[string]string
}

type RequestOptions struct {
	HttpMethod string
	Query      map[string]string
	Headers    map[string]string
	Body       interface{}
}

func NewRestApiClient(youtrackUrl string, token string) (*Client, error) {
	resolvedUrl, err := url.Parse(youtrackUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		YoutrackUrl: resolvedUrl,
		Token:       token,
		DefaultHeaders: map[string]string{
			"Accept":        "application/json",
			"Cache-Control": "no-cache",
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		},
	}, nil
}

func (client *Client) Fetch(resourceUrl string, requestParams *RequestOptions, result interface{}) error {
	resolvedUrl, err := url.Parse(resourceUrl)
	if err != nil {
		return err
	}

	// Append query params
	if requestParams.Query != nil {
		queryValues, _ := url.ParseQuery(resolvedUrl.RawQuery)
		for k, v := range requestParams.Query {
			queryValues.Add(k, v)
		}
		resolvedUrl.RawQuery = queryValues.Encode()
	}

	requestUrl := client.YoutrackUrl.ResolveReference(resolvedUrl)

	requestBody := new(bytes.Buffer)
	if requestParams.Body != nil {
		encoder := json.NewEncoder(requestBody)
		if err := encoder.Encode(requestParams.Body); err != nil {
			return err
		}
	}

	request, err := http.NewRequest(requestParams.HttpMethod, requestUrl.String(), requestBody)
	if err != nil {
		return err
	}

	// add default headers
	for k, v := range client.DefaultHeaders {
		request.Header.Set(k, v)
	}

	// add additional headers
	if requestParams.Headers != nil {
		for k, v := range requestParams.Headers {
			request.Header.Set(k, v)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.FromError(err, response.StatusCode, requestParams.HttpMethod, resourceUrl)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return errors.New(string(bodyBytes), response.StatusCode, requestParams.HttpMethod, resourceUrl)
	}

	if result != nil {
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(result)
		if err != nil {
			return err
		}
	}

	return nil
}

func (client *Client) Get(resourceUrl string, headers map[string]string, result interface{}) error {
	return client.Fetch(resourceUrl, &RequestOptions{
		HttpMethod: "GET",
		Headers:    headers,
		Body:       nil,
	}, result)
}

func (client *Client) Post(resourceUrl string, request interface{}, headers map[string]string, result interface{}) error {
	return client.Fetch(resourceUrl, &RequestOptions{
		HttpMethod: "POST",
		Headers:    headers,
		Body:       request,
	}, result)
}
