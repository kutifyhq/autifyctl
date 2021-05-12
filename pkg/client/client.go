package client

import (
	"errors"
	"net/http"

	"github.com/kutifyhq/go-autify/autify"
)

const (
	// Derived from here
	// https://autifyhq.github.io/autify-api/
	defaultAPIURL = "https://app.autify.com/api/v1/"
)

func NewClient(accesToken string) (autify.ClientWithResponsesInterface, error) {
	if accesToken == "" {
		return nil, errors.New("access token must be passed")
	}

	httpClient := http.DefaultClient

	opts := []autify.ClientOption{
		autify.WithHTTPClient(httpClient),
	}

	client, err := autify.NewClientWithResponses(defaultAPIURL, opts...)
	if err != nil {
		return nil, err
	}

	return client, err
}
