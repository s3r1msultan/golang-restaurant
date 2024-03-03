package routers

import (
	"final/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func AdminRouter(adminRouter *mux.Router) {
	adminRouter.StrictSlash(true)
	adminRouter.HandleFunc("/", controllers.AdminPageHandler)
	adminRouter.HandleFunc("/send_message", controllers.NotifyVerifiedUsers).Methods(http.MethodPost)
	adminRouter.HandleFunc("/api/users/{id}", controllers.GetUser).Methods(http.MethodGet)
	adminRouter.HandleFunc("/api/users", controllers.ListUsers).Methods(http.MethodGet)
	adminRouter.HandleFunc("/api/users", controllers.AddUser).Methods(http.MethodPost)
	adminRouter.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	adminRouter.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
}
