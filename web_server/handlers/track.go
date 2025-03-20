package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/providers/spotify"
	"github.com/kingtingthegreat/top-fetch/web_server/db"
)

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("no id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no id provided"))
		return
	}

	choice, err := strconv.Atoi(r.URL.Query().Get("choice"))
	if err != nil {
		choice = 1
	}

	user, err := db.GetUserById(id)
	if err != nil {
		log.Println("no user found", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id. user not found."))
		return
	}

	track, newAccessToken, err := spotify.GetUserTopTrack(env.EnvVal("SPOTIFY_CLIENT_ID"), env.EnvVal("SPOTIFY_CLIENT_SECRET"), user.AccessToken, user.RefreshToken, choice)
	if err != nil {
		log.Println("could not get track")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong. please try again."))
		return
	}

	if newAccessToken != "" {
		user.AccessToken = newAccessToken
		db.UpdateUser(user)
	}

	log.Println("song:", track.Name, "by", track.Artists[0].Name, "on", track.Album.Name)
	link := fmt.Sprintf(" \x1B]8;;%s\x1B\\🎵\x1B]8;;\x1B\\", track.ExternalUrls.Spotify)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(track.Album.Images[0].Url + "\x1d" + track.Name + " - " + track.Artists[0].Name + link))
}
