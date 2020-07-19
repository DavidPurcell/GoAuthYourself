package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	model "github.com/davidpurcell/GoAuthYourself/app/models"
	"go.mongodb.org/mongo-driver/bson"
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

	collection := client.Database("user").Collection("authorization")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.M{"Id": user.Id, "Username": user.Username, "Password": user.Password})
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, res)
}
