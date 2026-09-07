package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port   string
	DBPath string
}

func loadConfig() Config {
	return Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DATABASE_URL", "./data/orders.db"),
	}
}


func getEnv(key,defaultValue string) string {
	if value :=os.Getenv(key);value !=""{
		return value
	}

	return defaultValue
}

