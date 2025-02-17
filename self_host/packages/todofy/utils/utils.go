package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

// ParseAllowedUsers parses a comma-separated list of allowed users in the format "username:password"
func ParseAllowedUsers(users string) (map[string]string, string) {
	parsedUsers := make(map[string]string)
	parsedUserStrings := ""
	for _, user := range strings.Split(users, ",") {
		parts := strings.Split(user, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid user format: %s. Expected 'username:password'", user)
		}
		parsedUsers[parts[0]] = parts[1]
		parsedUserStrings += fmt.Sprintf("%s:%s, ", parts[0], "<hidden>")
	}
	parsedUserStrings = strings.TrimSuffix(parsedUserStrings, ", ")
	return parsedUsers, parsedUserStrings
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
