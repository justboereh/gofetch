package gofetch

import (
	"fmt"
	"net/http"
)

type FetchOptions struct {
	BaseURL       string
	Method        string
	Headers       map[string]string
	ResponseType  string
	ParseResponse func(response string) interface{}
	Query         map[string]interface{}
	Body          interface{}
	Retry         int
	RetryDelay    int
	Timeout       int
}

type FetchReponse struct{}

func Fetch(path string, options FetchOptions) FetchReponse {
	client := &http.Client{}

	fmt.Println("hello world")

	return FetchReponse{}
}
