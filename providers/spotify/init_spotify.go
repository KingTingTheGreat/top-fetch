package spotify

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func InitSpotify(clientId, clientSecret string) (string, string, error) {
	var accessToken, refreshToken string
	stop := make(chan bool, 1)

	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			code := r.URL.Query().Get("code")
			if code == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("no authorization code"))
				return
			}

			var err error
			accessToken, refreshToken, err = ExchangeCode(
				clientId,
				clientSecret,
				"http://localhost:8080/callback/spotify",
				code,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something went wrong. please try again."))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success! you can now close this window."))

			go func() {
				time.Sleep(100 * time.Millisecond)
				stop <- true
			}()
		}),
	}

	go func() {
		server.ListenAndServe()
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return "", "", fmt.Errorf("faild to gracefully shutdown server")
	}

	if accessToken == "" || refreshToken == "" {
		return "", "", fmt.Errorf("failed to get access token and/or refresh token")
	}

	return accessToken, refreshToken, nil
}
