package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch/web_server/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range os.Environ() {
		fmt.Println(k, v)
	}
	server := server.Server()
	server.Handler.ServeHTTP(w, r)
}
