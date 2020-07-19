package handler

import (
	"encoding/json"
	"net/http"

	model "github.com/davidpurcell/GoAuthYourself/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "DATA GOES HERE")
}

func CreateUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	respondJSON(w, http.StatusCreated, user)
}
