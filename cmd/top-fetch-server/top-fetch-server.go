package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kingtingthegreat/top-fetch/web_server/server"
)

func main() {
	godotenv.Load()

	server := server.Server()
	fmt.Println("Server running at http://localhost:8080")
	err := server.ListenAndServe()
	log.Fatal(err)
}
