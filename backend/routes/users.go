package routes

import (
	"ShaderLand/db"
	"ShaderLand/db/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const baseRoute = "/users"

func InitUsersRoutes(r *mux.Router) {
	r.HandleFunc(baseRoute, handleUserPost).Methods("POST", "OPTIONS")
	r.HandleFunc(baseRoute, handleUsersGet).Methods("GET")
	r.HandleFunc(baseRoute+"/{id}", handleUserDelete).Methods("DELETE")
	r.HandleFunc(baseRoute+"/password", handleUserPasswordCheck).Methods("POST", "OPTIONS")
}

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var userReq models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	user, err := db.DbRepo.Users.Create(userReq)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
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

func handleUsersGet(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	userName := r.URL.Query().Get("userName")
	if userName != "" {
		handleUserGet(w, userName)
		return
	}

	users, err := db.DbRepo.Users.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(users)
}

func handleUserGet(w http.ResponseWriter, userName string) {
	user, err := db.DbRepo.Users.GetByUserName(userName)
	if err == nil {
		json.NewEncoder(w).Encode(user)
		return
	}

	if errors.Is(err, models.ErrNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

func handleUserDelete(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	err = db.DbRepo.Users.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusAccepted)
}

func handleUserPasswordCheck(w http.ResponseWriter, r *http.Request) {
	addResponseHeaders(&w)
	if r.Method == "OPTIONS" {
		return
	}

	query := r.URL.Query()
	password := query.Get("password")
	userName := query.Get("userName")
	userRequest := models.UserRequest{
		UserName: userName,
		Password: password,
	}
	isValidUser := db.DbRepo.Users.IsValidPassword(userRequest)

	response := map[string]bool{
		"isValid": isValidUser,
	}
	json.NewEncoder(w).Encode(response)
}
