package main

import (
	"ShaderLand/sheets"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var env map[string]string

func main() {
	temp, err := godotenv.Read()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	env = temp

	r := mux.NewRouter()
	r.HandleFunc("/test", handleTest).Methods("GET")
	r.HandleFunc("/test", handleTestPost).Methods("POST")

	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")

	response, err := sheets.GetShaderDetail(env["SPREAD_SHEET_ID"], "test")
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func handleTestPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	data := sheets.ShaderDetail{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Bad Request", "detail": "Request body not correctly structured"}
		json.NewEncoder(w).Encode(response)
		return
	}

	err = sheets.AppendShaderDetail(env["SPREAD_SHEET_ID"], data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"message": "Internal Server Error", "detail": fmt.Sprint(err)}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
