package fetch

import (
	"fmt"
	"log"

	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/providers/spotify"
)

func LocalFetch() (string, string, error) {
	cfg := config.Config()
	if cfg.SpotifyClientId == "" || cfg.SpotifyClientSecret == "" {
		log.Fatal("Spotify client id or client secret is not set")
	}
	if cfg.SpotifyAccessToken == "" || cfg.SpotifyRefreshToken == "" {
		fmt.Println("please visit", spotify.AuthUrl(cfg.SpotifyClientId, "http://localhost:8080/callback/spotify"))
		var err error
		cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, err = spotify.InitSpotify(cfg.SpotifyClientId, cfg.SpotifyClientSecret)
		if err != nil {
			log.Fatal(err)
		}

		if cfg.Env == "" {
			go env.SaveSpotifyEnv(cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
		} else {
			go env.SaveSpotifyEnvFile(cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, cfg.Env)
		}
	}
	// log.Println("getting top track")
	track, newAccessToken, err := spotify.GetUserTopTrack(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, cfg.Choice)
	if err != nil {
		return "", "", err
	}
	if newAccessToken != "" {
		if cfg.Env == "" {
			go env.SaveSpotifyEnv(cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
		} else {
			go env.SaveSpotifyEnvFile(cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, cfg.Env)
		}
	}

	link := fmt.Sprintf(" \x1B]8;;%s\x1B\\🎵\x1B]8;;\x1B\\", track.ExternalUrls.Spotify)
	return track.Album.Images[0].Url, track.Name + " - " + track.Artists[0].Name + link, nil
}
