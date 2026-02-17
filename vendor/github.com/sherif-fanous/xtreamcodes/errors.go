package xtreamcodes

import (
	"fmt"
)

// HTTPError represents an error returned by the Xtream Codes API
type HTTPError struct {
	URL        string
	StatusCode int
	Body       *string
}

// Error implements the error interface
func (e *HTTPError) Error() string {
	if e.Body == nil {
		return fmt.Sprintf("HTTP request to %s failed with status code: %d", e.URL, e.StatusCode)
	}
	return fmt.Sprintf(
		"HTTP request to %s failed with status code: %d: body: %s",
		e.URL,
		e.StatusCode,
		*e.Body,
	)
}

// DecoderError represents an error that occurs while decoding API responses.
type DecoderError struct {
	URL                    string
	UnderlyingDecoderError error
}

// Error implements the error interface
func (e *DecoderError) Error() string {
	return fmt.Sprintf("failed to decode response from %s: %v", e.URL, e.UnderlyingDecoderError)
}

// Unwrap implements the Wrapper interface
func (e *DecoderError) Unwrap() error {
	return e.UnderlyingDecoderError
}
