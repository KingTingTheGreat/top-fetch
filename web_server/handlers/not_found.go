package handlers

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/web_server/tmplts"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmplts.LayoutString("Page Not Found", "404").Render(r.Context(), w)
}
