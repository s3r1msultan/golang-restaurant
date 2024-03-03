package routers

import (
	"final/controllers"
	"final/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthRouter(authRouter *mux.Router) {
	authRouter.StrictSlash(true)
	authRouter.Use(middlewares.FormToJSONMiddleware)
	authRouter.HandleFunc("/sign_up", controllers.SignupHandler).Methods(http.MethodPost, http.MethodGet)
	authRouter.HandleFunc("/sign_in", controllers.SigninHandler).Methods(http.MethodPost, http.MethodGet)
}
