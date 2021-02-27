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

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	serviceUrl      = "http://localhost:8001"
	secretKey       = "k8@0y%m^4-)ltn%8frs&e6^%dus1)6%s3&_u436h04)hjd6v#o"
	defaultPageSize = 10
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "ads"
)

const (
	enableHeader    = "Access-Control-Expose-Headers"
	jwtHeader       = "Authorization"
	firstPageHeader = "First-Page"
	lastPageHeader  = "Last-Page"
)

var db *gorm.DB = nil
var err error = nil

type Advertisment struct {
	ID          int
	CreatedOn   string
	UserId      int
	Email       string
	Name        string
	Category    string
	Price       int
	Description string
	Images      []Image
}

type Image struct {
	ID      int
	Path    string
	ProdRef int
}

func openDatabase() {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open("postgres", info)
	if err != nil {
		panic(err)
	}
}

func parseJWT(request *http.Request) jwt.MapClaims {
	token := request.Header.Get(jwtHeader)[4:]
	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return claims
}

func allAds(response http.ResponseWriter, request *http.Request) {
	openDatabase()
	defer db.Close()
	var ads []Advertisment
	var count int

	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	size, _ := strconv.Atoi(request.URL.Query().Get("size"))
	if size == 0 {
		size = defaultPageSize
	}
	search := "%" + strings.ToLower(request.URL.Query().Get("search")) + "%"

	db.Model(&Advertisment{}).Offset(page*size).Limit(size).Where("lower(name) like ? or lower(category) like ? or lower(description) like ?", search, search, search).Find(&ads)
	db.Model(&Advertisment{}).Where("lower(name) like ? or lower(category) like ? or lower(description) like ?", search, search, search).Count(&count)
	for index, product := range ads {
		db.Model(&Image{}).Where("prod_ref = ?", product.ID).Find(&ads[index].Images)
	}

	response.Header().Set(enableHeader, firstPageHeader+", "+lastPageHeader)
	response.Header().Set(firstPageHeader, strconv.FormatBool(page == 0))
	response.Header().Set(lastPageHeader, strconv.FormatBool(size*(page+1) >= count))

	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(ads)
}

func oneAd(response http.ResponseWriter, request *http.Request) {
	//a ako je prazan token??
	openDatabase()
	defer db.Close()
	var ad Advertisment
	var count int

	db.Model(&Advertisment{}).Where("id = ?", mux.Vars(request)["id"]).Find(&ad).Count(&count)
	if count == 0 {
		response.WriteHeader(404)
	}

	db.Model(&Image{}).Where("prod_ref = ?", ad.ID).Find(&ad.Images)
	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(ad)
}

func createAd(response http.ResponseWriter, request *http.Request) {
	//400 validacija
	//datum
	claims := parseJWT(request)
	openDatabase()
	defer db.Close()
	var ad Advertisment

	json.NewDecoder(request.Body).Decode(&ad)
	ad.UserId = claims["user_id"].(int)
	ad.Email = claims["email"].(string)
	db.Create(&ad)

	for _, image := range ad.Images {
		image.ProdRef = ad.ID
		var count int
		db.Model(&Image{}).Count(&count)
		data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
		path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
		ioutil.WriteFile(path, data, 0644)
		image.Path = serviceUrl + "/" + path
		db.Create(&image)
	}

	db.Model(&Image{}).Where("prod_ref = ?", ad.ID).Find(&ad.Images)
	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(ad)
}

func updateAd(response http.ResponseWriter, request *http.Request) {
	claims := parseJWT(request)
	openDatabase()
	defer db.Close()
	var ad Advertisment

	db.Where("id = ?", mux.Vars(request)["id"]).Find(&ad)
	json.NewDecoder(request.Body).Decode(&ad)
	if claims["user_id"].(int) != ad.UserId {
		response.WriteHeader(403)
		return
	}

	ad.UserId = claims["user_id"].(int)
	ad.Email = claims["email"].(string)
	db.Save(&ad)

	var images []Image
	db.Model(&Image{}).Where("prod_ref = ?", ad.ID).Find(&images)
	db.Model(&Image{}).Delete(&images)

	for _, image := range ad.Images {
		image.ProdRef = ad.ID
		if image.ID == 0 {
			var count int
			db.Model(&Image{}).Count(&count)
			data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
			path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
			ioutil.WriteFile(path, data, 0644)
			image.Path = serviceUrl + "/" + path
		}
		db.Create(&image)
	}

	db.Model(&Image{}).Where("prod_ref = ?", ad.ID).Find(&ad.Images)
	enc := json.NewEncoder(response)
	enc.SetIndent("", "    ")
	enc.Encode(ad)
}

func deleteAd(response http.ResponseWriter, request *http.Request) {
	claims := parseJWT(request)
	openDatabase()
	defer db.Close()
	var ad Advertisment
	var count int

	db.Model(&Advertisment{}).Where("id = ?", mux.Vars(request)["id"]).Find(&ad).Count(&count)
	if count == 0 {
		response.WriteHeader(404)
	}
	if claims["user_id"].(int) != ad.UserId {
		response.WriteHeader(403)
		return
	}
	db.Delete(&ad)
}

func databaseInit() {
	openDatabase()
	defer db.Close()
	//db.DropTableIfExists("ads")
	//db.DropTableIfExists("images")
	db.AutoMigrate(&Advertisment{})
	db.AutoMigrate(&Image{})
}

func routerInit() {
	router := mux.NewRouter()
	router.HandleFunc("/api/ads", allAds).Methods("GET")
	router.HandleFunc("/api/ads/{id}", allAds).Methods("GET")
	router.HandleFunc("/api/ads", allAds).Methods("POST")
	router.HandleFunc("/api/ads/{id}", allAds).Methods("PUT")
	router.HandleFunc("/api/ads/{id}", allAds).Methods("DELETE")

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
	databaseInit()
	routerInit()
}
