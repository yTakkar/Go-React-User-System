package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Route for all files
func Route(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplates(w, "app", &Page{"App", nil, nil})
}
