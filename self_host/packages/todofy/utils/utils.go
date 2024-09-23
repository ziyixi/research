package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func ExtractEnv(c *gin.Context) (map[string]string, error) {
	envs, ok := c.Get("envs")
	if !ok {
		return nil, fmt.Errorf("envs is not in context")
	}
	envsMap, ok := envs.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("envs is not a map")
	}
	return envsMap, nil
}

// FetchWithBasicAuth makes an HTTP GET request with Basic Auth and returns a dynamic JSON structure
func FetchWithBasicAuth(url, username, password string) (interface{}, error) {
	client := resty.New()

	// Make the request with Basic Auth
	resp, err := client.R().
		SetBasicAuth(username, password).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("error making HTTP request to %s: %w", url, err)
	}

	// Create a variable to hold the dynamic JSON response
	var result interface{}

	// Unmarshal the response body into the dynamic structure
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Return the parsed JSON as a generic interface{}
	return result, nil
}
