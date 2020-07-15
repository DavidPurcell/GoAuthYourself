package handler

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "DATA GOES HERE")
}
