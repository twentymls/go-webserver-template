package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey returns the API key from the request headers
// Example:
// Authorization: ApiKey {api_key}
func GetApiKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")

	if value == "" {
		return "", errors.New("no API key provided")
	}

	values := strings.Split(value, " ")

	// as specified, the API key should be in the format "ApiKey {api_key}"
	// so we expect 2 values if we split on the space
	if len(values) != 2 {
		return "", errors.New("invalid API key format")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("invalid API key format")
	}

	return values[1], nil
}
