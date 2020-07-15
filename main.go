package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(Users)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
	for _, user := range Users {
		if user.Id == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/getuserbyid/{id}", getUserById)

	myRouter.HandleFunc("/user", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Users = []User{
		User{Id: "1", Username: "User 1", Password: "YouShouldHashThis"},
		User{Id: "2", Username: "User 2", Password: "SeriouslyHashThese"},
	}
	handleRequests()
}

type User struct {
	Id       string
	Username string
	Password string
}

var Users []User
