package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Environment Variables:")
		for i, env := range os.Environ() {
			log.Println(i, env)
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
