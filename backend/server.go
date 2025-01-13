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
	r.HandleFunc("/shaders", handleShadersPost).Methods("POST", "OPTIONS")
	r.HandleFunc("/shaders/{name}", handleShaderGet).Methods("GET")

	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleShaderGet(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	shaderName := mux.Vars(r)["name"]
	response, err := sheets.GetShaderDetail(env[SPREAD_SHEET_ID], shaderName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func handleShadersPost(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	if r.Method == "OPTIONS" {
		return
	}

	data := sheets.ShaderDetailRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Bad Request", "detail": "Request body not correctly structured"}
		json.NewEncoder(w).Encode(response)
		return
	}

	sheetId, _ := strconv.Atoi(env[SHADERS_SHEET_ID])
	response, err := sheets.AppendShaderDetail(env[SPREAD_SHEET_ID], sheetId, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"message": "Internal Server Error", "detail": fmt.Sprint(err)}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func addResponseHeaders(w *http.ResponseWriter) {
	// CORS Stuff
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	(*w).Header().Add("Content-Type", "application/json")
}

func handleShadersGet(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	params := r.URL.Query()
	var startRow int = 2
	var endRow int = 20

	if params.Get("start") != "" {
		val, err := strconv.Atoi(params.Get("start"))
		if err == nil {
			startRow = val
		}
	}
	if params.Get("end") != "" {
		val, err := strconv.Atoi(params.Get("end"))
		if err == nil {
			endRow = val
		}
	}

	response, err := sheets.GetShaderDetails(env[SPREAD_SHEET_ID], startRow, endRow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"message": "Internal Server Error", "detail": fmt.Sprint(err)}
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(response)
}
