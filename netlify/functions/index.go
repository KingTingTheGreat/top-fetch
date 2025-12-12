package handler

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/web_server/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := server.Server()
	server.Handler.ServeHTTP(w, r)
}
