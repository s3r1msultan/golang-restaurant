package routers

import (
	"final/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SupportRouter(r *mux.Router) {
	r.HandleFunc("", controllers.SupportPageHandler).Methods(http.MethodGet, http.MethodPost)
}
