package handlers

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/tmplts"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	tmplts.LayoutComponent(tmplts.Home(), "Top Fetch").Render(r.Context(), w)
}
