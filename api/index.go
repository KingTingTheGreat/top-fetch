package handler

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/web_server/db"
	"github.com/kingtingthegreat/top-fetch/web_server/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db.DB = db.ConnectDB()
	db.UserCollection = db.GetCollection(env.EnvVal("COLLECTION_NAME"))
	server := server.Server()
	server.Handler.ServeHTTP(w, r)
}
