package fetch

import (
	"fmt"
	"log"

	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/providers/spotify"
)

func LocalFetch() (string, string) {
	cfg := config.Config()
	if cfg.SpotifyClientId == "" || cfg.SpotifyClientSecret == "" {
		log.Fatal("Spotify client id or client secret is not set")
	}
	if cfg.SpotifyAccessToken == "" || cfg.SpotifyRefreshToken == "" {
		fmt.Println("please visit", spotify.AuthUrl(cfg.SpotifyClientId, "http://localhost:8080"))
		var err error
		cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, err = spotify.InitSpotify(cfg.SpotifyClientId, cfg.SpotifyClientSecret)
		if err != nil {
			log.Fatal(err)
		}
		go env.SaveSpotifyEnv(cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	}
	// log.Println("getting top track")
	track, newAccessToken, err := spotify.GetUserTopTrack(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	if newAccessToken != "" {
		go env.SaveSpotifyEnv(newAccessToken, cfg.SpotifyRefreshToken)
	}

	return track.Album.Images[0].Url, track.Name + " - " + track.Artists[0].Name
}
