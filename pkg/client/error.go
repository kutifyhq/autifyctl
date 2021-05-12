package client

import (
	"encoding/json"
)

type APIError struct {
	Errors []Error
}

type Error struct {
	Message string
}

func ParseError(body []byte) (*APIError, error) {
	var apiError APIError

	if err := json.Unmarshal(body, &apiError); err != nil {
		return nil, err
	}

	return &apiError, nil
}
