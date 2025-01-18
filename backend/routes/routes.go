package routes

import "github.com/gorilla/mux"

func InitRoutes(r *mux.Router) {
	initShadersRoutes(r)
	InitUsersRoutes(r)
}
