package handlers

import (
	"log"
	"net/http"

	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/providers"
	"github.com/kingtingthegreat/top-fetch/providers/spotify"
	"github.com/kingtingthegreat/top-fetch/web_server/db"
	"github.com/kingtingthegreat/top-fetch/web_server/tmplts"
)

func SpotifyCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}

	clientId := env.EnvVal("SPOTIFY_CLIENT_ID")
	clientSecret := env.EnvVal("SPOTIFY_CLIENT_SECRET")
	redirectUri := env.EnvVal("SPOTIFY_REDIRECT_URI")

	accessToken, refreshToken, err := spotify.ExchangeCode(clientId, clientSecret, redirectUri, code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}

	spotifyId, newAccessToken, err := spotify.GetSpotifyId(clientId, clientSecret, accessToken, refreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}
	if newAccessToken != "" {
		accessToken = newAccessToken
	}

	// recover existing id
	user, err := db.GetUserByProviderId(spotifyId)
	if err == nil {
		log.Println("found existing user")
		w.WriteHeader(http.StatusOK)
		tmplts.LayoutComponent(tmplts.Callback(user.Id), "TopFetch").Render(r.Context(), w)
		return
	}

	// new user, create a new id for them
	user = &db.DBUser{
		ProviderId:   spotifyId,
		Provider:     providers.SPOTIFY,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	id, err := db.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}

	log.Println("created new user")
	tmplts.LayoutComponent(tmplts.Callback(id), "TopFetch").Render(r.Context(), w)
}
