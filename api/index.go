package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch/web_server/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Environment Variables:")
	for i, env := range os.Environ() {
		fmt.Println(i, env)
	}
	server := server.Server()
	server.Handler.ServeHTTP(w, r)
}
