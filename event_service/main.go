package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *gorm.DB = nil
var router *mux.Router = nil

func initRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/api/events", getEventsEndpoint).Methods("GET")
	router.HandleFunc("/api/events", createEventEndpoint).Methods("POST")
	router.HandleFunc("/api/events/{id:[0-9]+}", updateEventEndpoint).Methods("PUT")
	router.HandleFunc("/api/events/{id:[0-9]+}", deleteEventEndpoint).Methods("DELETE")
	router.HandleFunc("/api/statistic/{start:[0-9]+}/{end:[0-9]+}", statisticEndpoint).Methods("GET")

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	staticDir := "/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	server := http.Server{
		Addr:    ":8002",
		Handler: cors.Handler(router),
	}

	log.Fatal(server.ListenAndServe())
}

func main() {
	initDatabase()
	initRouter()
}
