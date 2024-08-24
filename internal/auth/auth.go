package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the api key from the request headers
// Example:
// Authorization: ApiKey <api_key>
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no auth information provided")
	}

	apiKeys := strings.Split(apiKey, " ")
	if len(apiKeys) != 2 {
		return "", errors.New("invalid auth information provided")
	}

	if apiKeys[0] != "ApiKey" {
		return "", errors.New("malformed auth information provided")
	}

	return apiKeys[1], nil
}
