package main

import (
	"ShaderLand/sheets"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var env map[string]string
var SPREAD_SHEET_ID = "SHEETS_SPREAD_SHEET_ID"
var SHADERS_SHEET_ID = "SHEETS_SHADERS_SHEET_ID"

func main() {
	temp, err := godotenv.Read(".env", ".env.local")
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	env = temp

	r := mux.NewRouter()
	r.HandleFunc("/shaders", handleShadersGet).Methods("GET")
	r.HandleFunc("/shaders", handleShadersPost).Methods("POST")

	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleShadersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")

	shaderName := r.URL.Query().Get("name")
	response, err := sheets.GetShaderDetail(env[SPREAD_SHEET_ID], shaderName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func handleShadersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	data := sheets.ShaderDetail{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Bad Request", "detail": "Request body not correctly structured"}
		json.NewEncoder(w).Encode(response)
		return
	}

	sheetId, _ := strconv.Atoi(env[SHADERS_SHEET_ID])
	err = sheets.AppendShaderDetail(env[SPREAD_SHEET_ID], int64(sheetId), data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"message": "Internal Server Error", "detail": fmt.Sprint(err)}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
