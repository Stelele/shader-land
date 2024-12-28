package main

import (
	"ShaderLand/sheets"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", handleTest).Methods("GET")

	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	response := sheets.GetSheetData()

	w.Header().Add("Content-Type", "text/json")
	json.NewEncoder(w).Encode(response)
}
