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
	router.HandleFunc("/api/ads", getAdsEndpoint).Methods("GET")
	router.HandleFunc("/api/ads-my", getMyAdsEndpoint).Methods("GET")
	router.HandleFunc("/api/ads/{id:[0-9]+}", getAdEndpoint).Methods("GET")
	router.HandleFunc("/api/ads", createAdEndpoint).Methods("POST")
	router.HandleFunc("/api/ads/{id:[0-9]+}", updateAdEndpoint).Methods("PUT")
	router.HandleFunc("/api/ads/{id:[0-9]+}", deleteAdEndpoint).Methods("DELETE")
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
		Addr:    ":8001",
		Handler: cors.Handler(router),
	}

	log.Fatal(server.ListenAndServe())
}

func main() {
	initDatabase()
	initRouter()
}
