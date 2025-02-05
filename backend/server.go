package main

import (
	"ShaderLand/db"
	"ShaderLand/routes"
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env", ".env.local")
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	clerk.SetKey(env["CLERK_SECRET_KEY"])

	db.InitDb()

	r := mux.NewRouter()
	routes.InitRoutes(r)

	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
