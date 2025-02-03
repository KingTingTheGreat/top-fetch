package env

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

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

func SaveSpotifyEnv(accessToken string, refreshToken string) error {
	basePath, err := GetBasePath()
	if err != nil {
		return err
	}

	env["SPOTIFY_ACCESS_TOKEN"] = accessToken
	env["SPOTIFY_REFRESH_TOKEN"] = refreshToken
	return godotenv.Write(env, filepath.Join(basePath, "env/.env"))
}

func EnvVal(key string) string {
	val := env[key]
	if val != "" {
		return val
	}
	return os.Getenv(key)
}
