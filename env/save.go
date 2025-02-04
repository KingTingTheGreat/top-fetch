package env

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

func SaveSpotifyEnv(accessToken string, refreshToken string) error {
	basePath, err := GetBasePath()
	if err != nil {
		return err
	}
	envFile := filepath.Join(basePath, "env/.env")

	return SaveSpotifyEnvFile(accessToken, refreshToken, envFile)
}

func SaveSpotifyEnvFile(accessToken string, refreshToken string, envFile string) error {
	env["SPOTIFY_ACCESS_TOKEN"] = accessToken
	env["SPOTIFY_REFRESH_TOKEN"] = refreshToken

	return godotenv.Write(env, envFile)
}
