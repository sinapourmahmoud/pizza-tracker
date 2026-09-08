package main

import (
	"html/template"
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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func loadTemplates(router *gin.Engine) error {

	functions := template.FuncMap{

		"add": func(a, b int) int { return a + b },
	}

	templ,err :=template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")

	if err !=nil{
		return err
	}

	router.SetHTMLTemplate(templ)

	return nil

}
