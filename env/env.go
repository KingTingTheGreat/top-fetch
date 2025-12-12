package env

import (
	_ "embed"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//go:embed .env
var envString string
var env map[string]string

func LoadEnv() {
	var err error
	env, err = godotenv.Unmarshal(envString)
	if err != nil {
		log.Fatal("failed to parse .env in env/.env")
	}
}

func EnvVal(key string) string {
	val := env[key]
	if val != "" {
		log.Printf("env: using %s from embedded .env: %s\n", key, val)
		return val
	}
	log.Printf("env: using %s from system environment: %s\n", key, val)
	return os.Getenv(key)
}
