package env

import (
	"os"
)

var (
	Host      = GetEnv("HOST", "0.0.0.0")
	Port      = GetEnv("PORT", "8888")
	DBHost    = GetEnv("DB_HOST", "127.0.0.1")
	DBPort    = GetEnv("DB_PORT", "3306")
	DBUser    = GetEnv("DB_USER", "root")
	DBPass    = GetEnv("DB_PASS", "example")
	DBName    = GetEnv("DB_NAME", "database")
	TLSOption = GetEnv("TLS_OPTION", "true")
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
