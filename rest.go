package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	headers map[string]string
}

/*
A high level function thats create a innstace for HttpClient
*/
func New() HttpClient {
	return HttpClient{}
}

func (h *HttpClient) SetHeaders(headers map[string]string) error {
	h.headers = headers
	return errors.New("Set the headers is not possible")
}

/*
HttpClient.Get
params -> url string
return -> string, error

this function execute a GET http request for the url expecified in the params
*/
func (h HttpClient) Get(url string) (string, error) {
	c := http.Client{Timeout: time.Duration(0) * time.Second}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	for key, value := range h.headers {
		req.Header.Add(key, value)
	}

	resp, err := c.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (h HttpClient) Post(url string, data map[string]interface{}) (string, error) {
	dataJson, _ := json.Marshal(data)
	c := http.Client{Timeout: time.Duration(0) * time.Second}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataJson))
	if err != nil {
		return "", err
	}
	for key, value := range h.headers {
		req.Header.Add(key, value)
	}
	resp, err := c.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
