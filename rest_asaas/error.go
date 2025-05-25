package rest_asaas

import (
	"encoding/json"
	"errors"
)

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

func (e *ErrorResponse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *ErrorResponse) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ErrorResponse) String() string {
	if len(e.Errors) == 0 {
		return ""
	}
	result := ""
	for _, err := range e.Errors {
		result += "Code: " + err.Code + ", Description: " + err.Description + "\n"
	}
	return result
}

func (e *ErrorResponse) Return() error {
	if len(e.Errors) == 0 {
		return nil
	}
	result := errors.New(e.String())
	return result
}

func NewErrorResponse(data []byte) (*ErrorResponse, error) {
	response := &ErrorResponse{}
	if err := response.Unmarshal(data); err != nil {
		return nil, err
	}
	return response, nil
}
