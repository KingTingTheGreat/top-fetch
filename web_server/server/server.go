package server

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/web_server/middleware"
	"github.com/kingtingthegreat/top-fetch/web_server/router"
)

func Server() *http.Server {
	router := router.Router()

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(router),
	}

	return &server
}
