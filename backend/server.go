package main

import (
	"ShaderLand/db"
	"ShaderLand/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDb()

	r := mux.NewRouter()
	routes.InitRoutes(r)

	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
