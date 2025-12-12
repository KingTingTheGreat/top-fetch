package handler

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/web_server/db"
	"github.com/kingtingthegreat/top-fetch/web_server/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db.ConnectDB()
	server := server.Server()
	server.Handler.ServeHTTP(w, r)
}
