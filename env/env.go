package env

import (
	_ "embed"
	"log"

	"github.com/joho/godotenv"
)

//go:embed .env
var EnvString string
var Env map[string]string

func LoadEnv() {
	var err error
	Env, err = godotenv.Unmarshal(EnvString)
	if err != nil {
		log.Fatal("failed to parse .env in env/.env")
	}
}

func SaveSpotifyEnv(accessToken string, refreshToken string) error {
	Env["SPOTIFY_ACCESS_TOKEN"] = accessToken
	Env["SPOTIFY_REFRESH_TOKEN"] = refreshToken
	return godotenv.Write(Env, ".env")
}
