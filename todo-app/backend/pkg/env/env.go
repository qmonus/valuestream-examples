package env

import (
	"os"
)

var Host = GetEnv("HOST", "0.0.0.0")
var Port = GetEnv("PORT", "8888")
var DBHost = GetEnv("DB_HOST", "127.0.0.1")
var DBPort = GetEnv("DB_PORT", "3306")
var DBUser = GetEnv("DB_USER", "root")
var DBPass = GetEnv("DB_PASS", "example")
var DBName = GetEnv("DB_NAME", "database")

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
