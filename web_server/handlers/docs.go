package handlers

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/web_server/tmplts"
)

func DocumentationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tmplts.LayoutComponent(tmplts.Docs(), "Docs").Render(r.Context(), w)
}
