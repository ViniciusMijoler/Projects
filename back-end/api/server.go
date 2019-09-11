package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"projects/back-end/api/handler"

	"github.com/gorilla/mux"
)

// App struct ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//StartServer ...
func (a *App) StartServer() {
	a.Router = mux.NewRouter()
	s := a.Router.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)

	s.HandleFunc("/projects", handler.InsertProject).Methods(http.MethodPost)
	s.HandleFunc("/projects/{id:[0-9]+}", handler.UpdateProject).Methods(http.MethodPut)
	s.HandleFunc("/projects/{id:[0-9]+}", handler.DeleteProject).Methods(http.MethodDelete)
	s.HandleFunc("/projects/{id:[0-9]+}", handler.GetProject).Methods(http.MethodGet)
	s.HandleFunc("/projects", handler.GetProjects).Methods(http.MethodGet)

	a.Router.Handle("/api/v1/{_:.*}", a.Router)
	port := 10001
	log.Printf("Starting Server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}
