package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	serviceURL      = "http://localhost:8002"
	secretKey       = "k8@0y%m^4-)ltn%8frs&e6^%dus1)6%s3&_u436h04)hjd6v#o"
	defaultPageSize = 10
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "events"
)

const (
	enableHeader    = "Access-Control-Expose-Headers"
	jwtHeader       = "Authorization"
	firstPageHeader = "First-Page"
	lastPageHeader  = "Last-Page"
)

var db *gorm.DB = nil
var err error = nil

type Event struct {
	ID          int
	CreatedOn   string
	UserId      int
	Email       string
	ProductId   int
	Name        string
	Category    string
	From        string
	To          string
	Place       string
	Description string
	Images      []Image
}

type Image struct {
	ID       int
	Path     string
	EventRef int
}

func openDatabase() {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open("postgres", info)
	if err != nil {
		panic(err)
	}
}

func parseJWT(request *http.Request) jwt.MapClaims {
	if len(request.Header.Get(jwtHeader)) < 4 {
		return nil
	}
	token := request.Header.Get(jwtHeader)[4:]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil
	}

	return claims
}

func allEvents(response http.ResponseWriter, request *http.Request) {
	claims := parseJWT(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	openDatabase()
	defer db.Close()
	var events []Event
	var count int

	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	size, _ := strconv.Atoi(request.URL.Query().Get("size"))
	if size == 0 {
		size = 10
	}
	productID, _ := strconv.Atoi(request.URL.Query().Get("product"))

	db.Model(&Event{}).Offset(page*size).Limit(size).Where("product_id = ?", productID).
		Order("created_on desc").Find(&events)
	db.Model(&Event{}).Where("product_id = ?", productID).Count(&count)
	for index, event := range events {
		db.Model(&Image{}).Where("event_ref = ?", event.ID).Find(&events[index].Images)
	}

	response.Header().Set(enableHeader, firstPageHeader+", "+lastPageHeader)
	response.Header().Set(firstPageHeader, strconv.FormatBool(page == 0))
	response.Header().Set(lastPageHeader, strconv.FormatBool(size*(page+1) >= count))
	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(events)
}

func createEvent(response http.ResponseWriter, request *http.Request) {
	claims := parseJWT(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	openDatabase()
	defer db.Close()
	var event Event

	json.NewDecoder(request.Body).Decode(&event)
	event.CreatedOn = time.Now().UTC().String()
	event.UserId = int(claims["user_id"].(float64))
	event.Email = claims["email"].(string)
	db.Create(&event)

	for _, image := range event.Images {
		image.EventRef = event.ID
		var count int
		db.Model(&Image{}).Count(&count)
		data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
		path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
		ioutil.WriteFile(path, data, 0644)
		image.Path = serviceURL + "/" + path
		db.Create(&image)
	}

	db.Where("event_ref = ?", event.ID).Find(&event.Images)
	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(event)
}

func updateEvent(response http.ResponseWriter, request *http.Request) {
	claims := parseJWT(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	openDatabase()
	defer db.Close()
	var event Event

	db.Where("id = ?", mux.Vars(request)["id"]).Find(&event)
	json.NewDecoder(request.Body).Decode(&event)
	if int(claims["user_id"].(float64)) != event.UserId {
		response.WriteHeader(403)
		return
	}

	event.UserId = int(claims["user_id"].(float64))
	event.Email = claims["email"].(string)
	db.Save(&event)
	var images []Image
	db.Model(&Image{}).Where("event_ref = ?", event.ID).Find(&images)
	db.Model(&Image{}).Delete(&images)

	for _, image := range event.Images {
		image.EventRef = event.ID
		if image.ID == 0 {
			var count int
			db.Model(&Image{}).Count(&count)
			data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
			path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
			ioutil.WriteFile(path, data, 0644)
			image.Path = serviceURL + "/" + path
		}
		db.Create(&image)
	}

	db.Where("event_ref = ?", event.ID).Find(&event.Images)
	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(event)
}

func deleteEvent(response http.ResponseWriter, request *http.Request) {
	claims := parseJWT(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	openDatabase()
	defer db.Close()
	var event Event
	var count int

	db.Model(&Event{}).Where("id = ?", mux.Vars(request)["id"]).Find(&event).Count(&count)
	if count == 0 {
		response.WriteHeader(404)
	}
	if int(claims["user_id"].(float64)) != event.UserId {
		response.WriteHeader(403)
		return
	}
	db.Delete(&event)
}

func databaseInit() {
	openDatabase()
	defer db.Close()
	db.DropTableIfExists("events")
	db.DropTableIfExists("images")
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Image{})
}

func routerInit() {
	router := mux.NewRouter()
	router.HandleFunc("/api/events", allEvents).Methods("GET")
	router.HandleFunc("/api/events", createEvent).Methods("POST")
	router.HandleFunc("/api/events/{id}", updateEvent).Methods("PUT")
	router.HandleFunc("/api/events/{id}", deleteEvent).Methods("DELETE")

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
	databaseInit()
	routerInit()
}
