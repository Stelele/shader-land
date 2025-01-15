package routes

import (
	"ShaderLand/db"
	"ShaderLand/db/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

const baseUrl = "/shaders"

func initShadersRoutes(r *mux.Router) {
	r.HandleFunc(baseUrl+"/{url}", handleShaderGet).Methods("GET")
	r.HandleFunc(baseUrl, handleShadersGet).Methods("GET")
	r.HandleFunc(baseUrl, handleShadersPost).Methods("POST", "OPTIONS")
}

func handleShadersGet(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	name := r.URL.Query().Get("name")
	var all []models.Shader
	var err error

	if name != "" {
		all, err = db.DbRepo.Shaders.GetByName(name)
	} else {
		all, err = db.DbRepo.Shaders.All()
	}

	if err == nil {
		json.NewEncoder(w).Encode(all)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

func handleShaderGet(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	vars := mux.Vars(r)
	shader, err := db.DbRepo.Shaders.GetByUrl(vars["url"])
	if err == nil {
		json.NewEncoder(w).Encode(shader)
		return
	}

	if errors.Is(err, models.ErrNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

func handleShadersPost(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	if r.Method == "OPTIONS" {
		return
	}

	var req models.ShaderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	shader, err := db.DbRepo.Shaders.Create(req)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shader)
		return
	}

	if errors.Is(err, models.ErrDuplicate) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}
