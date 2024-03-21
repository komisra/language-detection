package langDetect

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"fmt"
)

// APIError is an error returned from the API
type APIError struct {
	// HTTP text status of response
	Status string
	// HTTP numerical status code of  response.
	StatusCode int
	// error code returned by API
	Code int `json:"code"`
	// human readable error message 
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.Status)
}

type apiErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

// detection fails --> DetectionError
type DetectionError struct {
	Message string
}

func (e *DetectionError) Error() string {
	return e.Message
}