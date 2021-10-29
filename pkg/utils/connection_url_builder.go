package utils

import (
	"fmt"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=true",
			GetEnv("DB_USER", ""),
			GetEnv("DB_PASSWORD", ""),
			GetEnv("DB_CONNECTION", ""),
			GetEnv("DB_HOST", ""),
			GetEnv("DB_PORT", ""),
			GetEnv("DB_NAME", ""),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			GetEnv("REDIS_HOST", ""),
			GetEnv("REDIS_PORT", ""),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			GetEnv("SERVER_HOST", ""),
			GetEnv("SERVER_PORT", ""),
		)
		fmt.Println(url)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
