package utils

import (
	"fmt"
	"log"
	"strings"
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
