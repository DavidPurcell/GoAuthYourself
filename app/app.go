package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidpurcell/GoAuthYourself/app/handler"
	"github.com/davidpurcell/GoAuthYourself/config"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	Mongo  *mongo.Client
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("mongodb+srv://admin:Password1!@goauthyourself.znhhb.mongodb.net/Authorization?retryWrites=true&w=majority")

	// Return to this later.
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/users", a.handleRequest(handler.GetAllUsers))
	a.Post("/user", a.handleRequest(handler.CreateUser))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *mongo.Client, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.Mongo, w, r)
	}
}
