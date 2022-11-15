package config

import (
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

func MustGet(key string, def string) string {
	value := os.Getenv(key)
	if def != "" && value == "" {
		return def
	} else if def == "" && value == "" {
		panic(key + ": not found in env")
	}
	return value
}

func init() {
	env := MustGet("ENV_PATH", ".env")
	log.Printf("Loading %s file\n", env)
	if er := godotenv.Load(env); er != nil {
		if er := godotenv.Load(); er != nil {
			log.Fatalf("Error loading %s file\n", env)
		}
	}
}
