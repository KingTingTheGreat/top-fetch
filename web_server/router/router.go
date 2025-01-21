package router

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/providers/spotify"
	"github.com/kingtingthegreat/top-fetch/web_server/handlers"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", handlers.HomePageHandler)
	router.HandleFunc("/docs", handlers.DocumentationHandler)

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.HandleFunc("GET /sign-in", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, spotify.AuthUrl(env.EnvVal("SPOTIFY_CLIENT_ID"), env.EnvVal("SPOTIFY_REDIRECT_URI")), http.StatusSeeOther)
	})
	router.HandleFunc("GET /callback/spotify", handlers.SpotifyCallbackHandler)

	router.HandleFunc("GET /track", handlers.TrackHandler)

	router.HandleFunc("/404", handlers.NotFoundHandler)

	return router
}
