package handlers

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/web_server/tmplts"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		tmplts.LayoutString("Page Not Found", "404").Render(r.Context(), w)
		return
	}

	tmplts.LayoutComponent(tmplts.Home(), "Top Fetch").Render(r.Context(), w)
}
