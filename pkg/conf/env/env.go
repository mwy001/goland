package env

import (
	"fmt"
	"os"
)

func CurrentEnvironment() string {
	result := getenvWithFallback("TASK_ENVIRONMENT", "local")
	fmt.Printf("TASK_ENVIRONMENT: %v\n", result)

	return result
}

func ElasticSearchLogEnabled() string {
	result := getenvWithFallback("ES_LOG_ENABLED", "0")
	fmt.Printf("ES_LOG_ENABLED: %v\n", result)

	return result
}

func LogstashLoggingEnabled() string {
	result := getenvWithFallback("LOGSTASH_LOGGING_ENABLED", "0")
	fmt.Printf("LOGSTASH_LOGGING_ENABLED: %v\n", result)

	return result
}

func WorkingDir(fallback string) string {
	result := getenvWithFallback("WORKING_DIR", fallback)
	fmt.Printf("WORKING_DIR: %v\n", result)

	return result
}

func getenvWithFallback(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
