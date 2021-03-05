package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func init() {
	initDatabase()
}

func TestGetAdsEmpty(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 0 {
		t.Error("advertisements list should be empty")
	}

}

func TestGetAdsWithLimit(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads?size=5", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 5 {
		t.Error("advertisements list should be length of 5")
	}

	order := ads[0].Name == "Advertisement 24" && ads[1].Name == "Advertisement 25" &&
		ads[2].Name == "Advertisement 26" && ads[3].Name == "Advertisement 27" && ads[4].Name == "Advertisement 28"
	if !order {
		t.Error("invalid advertisements data")
	}
}

func TestGetAdsWithOffset(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads?page=1&size=5", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 5 {
		t.Error("advertisements list should be length of 5")
	}

	order := ads[0].Name == "Advertisement 29" && ads[1].Name == "Advertisement 30" &&
		ads[2].Name == "Advertisement 21" && ads[3].Name == "Advertisement 22" && ads[4].Name == "Advertisement 23"
	if !order {
		t.Error("invalid advertisements data")
	}
}

func TestGetAdsWithSearch(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads?size=5&search=3", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 4 {
		t.Error("advertisements list should be length of 4")
	}

	order := ads[0].Name == "Advertisement 30" && ads[1].Name == "Advertisement 23" &&
		ads[2].Name == "Advertisement 13" && ads[3].Name == "Advertisement 3"
	if !order {
		t.Error("invalid advertisements data")
	}
}

func TestGetMyAdsEmpty(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads-my", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMyAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 0 {
		t.Error("advertisements list should be empty")
	}

}

func TestGetMyAdsWithLimit(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads-my?size=5", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMyAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 5 {
		t.Error("advertisements list should be length of 5")
	}

	order := ads[0].Name == "Advertisement 9" && ads[1].Name == "Advertisement 10" &&
		ads[2].Name == "Advertisement 6" && ads[3].Name == "Advertisement 7" && ads[4].Name == "Advertisement 8"
	if !order {
		t.Error("invalid advertisements data")
	}
}

func TestGetMyAdsWithOffset(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads-my?page=1&size=5", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMyAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 5 {
		t.Error("advertisements list should be length of 5")
	}

	order := ads[0].Name == "Advertisement 3" && ads[1].Name == "Advertisement 4" &&
		ads[2].Name == "Advertisement 5" && ads[3].Name == "Advertisement 1" && ads[4].Name == "Advertisement 2"
	if !order {
		t.Error("invalid advertisements data")
	}
}

func TestGetMyAdsWithSearch(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads-my?size=5&search=3", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMyAdsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ads []advertisement
	json.NewDecoder(rr.Body).Decode(&ads)
	if len(ads) != 1 {
		t.Error("advertisements list should be length of 1")
	}

	order := ads[0].Name == "Advertisement 3"
	if !order {
		t.Error("invalid advertisements data")
	}
}

func TestGetAdNotFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads/31", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", getAdEndpoint).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Error("status code should be 404 Not Found")
	}

}

func TestGetAd(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/ads/1", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", getAdEndpoint).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ad advertisement
	json.NewDecoder(rr.Body).Decode(&ad)

	order := ad.ID == 1 && ad.Active && ad.UserID == 1 &&
		ad.Name == "Advertisement 1" && ad.Category == "HONEY" && ad.Price == 123 &&
		ad.Description == "Description 1" && len(ad.Images) == 1 &&
		ad.Images[0].ID == 1 && ad.Images[0].Path == "http://localhost:8001/image.jpeg" &&
		ad.Images[0].ProdRef == 1
	if !order {
		t.Error("invalid advertisement data")
	}
}

func TestCreateAd(t *testing.T) {

	req, _ := http.NewRequest("POST", "/api/ads", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"Name": "test name",
		"Category": "test category",
		"Price": 123,
		"Description": "test description",
		"Images": [{
			"path": "%s"
		},{
			"path": "%s"
		}]
	}`, base64image(), base64image()))))
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createAdEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ad advertisement
	json.NewDecoder(rr.Body).Decode(&ad)

	order := ad.ID == 31 && ad.Active && ad.UserID == 1 &&
		ad.Name == "test name" && ad.Category == "test category" && ad.Price == 123 &&
		ad.Description == "test description" && len(ad.Images) == 2 &&
		ad.Images[0].ID == 31 && ad.Images[0].Path == "http://localhost:8001/image31.jpeg" && ad.Images[0].ProdRef == 31 &&
		ad.Images[1].ID == 32 && ad.Images[1].Path == "http://localhost:8001/image32.jpeg" && ad.Images[1].ProdRef == 31
	if !order {
		t.Error("invalid advertisement data")
	}

	if count := imagesNumber(); count != 32 {
		t.Error("invalid number of images")
	}

	if count := detachedImagesNumber(); count != 0 {
		t.Error("invalid number of detached images")
	}

	_, err1 := os.Stat("image31.jpeg")
	_, err2 := os.Stat("image32.jpeg")
	if err1 != nil || err2 != nil {
		t.Error("invalid file images")
	}

	os.Remove("image31.jpeg")
	os.Remove("image32.jpeg")
}

func TestCreateAdBadRequest(t *testing.T) {

	req, _ := http.NewRequest("POST", "/api/ads", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"Name": "",
		"Category": "",
		"Price": "",
		"Description": "",
		"Images": [{
			"path": "%s"
		},{
			"path": "%s"
		}]
	}`, base64image(), base64image()))))
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createAdEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("status code should be 400 Bad Request")
	}
}

func TestUpdateAd(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/ads/31", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"Name": "test name 2",
		"Category": "test category",
		"Price": 123,
		"Description": "test description",
		"Images": [{
			"path": "%s"
		},{
			"path": "%s"
		}]
	}`, base64image(), base64image()))))
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", updateAdEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ad advertisement
	json.NewDecoder(rr.Body).Decode(&ad)

	order := ad.ID == 31 && ad.Active && ad.UserID == 1 &&
		ad.Name == "test name 2" && ad.Category == "test category" && ad.Price == 123 &&
		ad.Description == "test description" && len(ad.Images) == 2 &&
		ad.Images[0].ID == 33 && ad.Images[0].Path == "http://localhost:8001/image33.jpeg" && ad.Images[0].ProdRef == 31 &&
		ad.Images[1].ID == 34 && ad.Images[1].Path == "http://localhost:8001/image34.jpeg" && ad.Images[1].ProdRef == 31
	if !order {
		t.Error("invalid advertisement data")
	}

	if count := imagesNumber(); count != 34 {
		t.Error("invalid number of images")
	}

	if count := detachedImagesNumber(); count != 2 {
		t.Error("invalid number of detached images")
	}

	_, err1 := os.Stat("image33.jpeg")
	_, err2 := os.Stat("image34.jpeg")
	if err1 != nil || err2 != nil {
		t.Error("invalid file images")
	}

	os.Remove("image33.jpeg")
	os.Remove("image34.jpeg")
}

func TestUpdateAdBadRequest(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/ads/1", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"Name": "",
		"Category": "",
		"Price": "",
		"Description": "",
		"Images": [{
			"path": "%s"
		},{
			"path": "%s"
		}]
	}`, base64image(), base64image()))))
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", updateAdEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("status code should be 400 Bad Request")
	}
}

func TestUpdateAdNotFound(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/ads/40", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"Name": "test name",
		"Category": "test category",
		"Price": 123,
		"Description": "test description",
		"Images": [{
			"path": "%s"
		},{
			"path": "%s"
		}]
	}`, base64image(), base64image()))))
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", updateAdEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Error("status code should be 404 Not Found")
	}
}

func TestUpdateAdForbidden(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/ads/11", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"Name": "test name",
		"Category": "test category",
		"Price": 123,
		"Description": "test description",
		"Images": [{
			"path": "%s"
		},{
			"path": "%s"
		}]
	}`, base64image(), base64image()))))
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", updateAdEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Error("status code should be 403 Forbidden")
	}
}

func TestDeleteAd(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/api/ads/31", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", deleteAdEndpoint).Methods("DELETE")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ad advertisement
	json.NewDecoder(rr.Body).Decode(&ad)

	order := !ad.Active
	if !order {
		t.Error("invalid advertisement data")
	}

	if count := imagesNumber(); count != 34 {
		t.Error("invalid number of images")
	}

	if count := detachedImagesNumber(); count != 4 {
		t.Error("invalid number of detached images")
	}
}

func TestDeleteAdNotFound(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/api/ads/40", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", deleteAdEndpoint).Methods("DELETE")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Error("status code should be 404 Not Found")
	}
}

func TestDeleteAdForbidden(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/api/ads/11", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/ads/{id:[0-9]+}", deleteAdEndpoint).Methods("DELETE")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Error("status code should be 403 Forbidden")
	}
}

func TestStatisticBadRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/statistic/2020/2020", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/statistic/{start:[0-9]+}/{end:[0-9]+}", statisticEndpoint).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("status code should be 400 Bad Request")
	}
}

func TestStatistic(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/statistic/2014/2020", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/statistic/{start:[0-9]+}/{end:[0-9]+}", statisticEndpoint).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	result := make([][2]int, 2020-2014+1)
	json.NewDecoder(rr.Body).Decode(&result)

	if fmt.Sprintf("%v", result) != "[[2014 2] [2015 3] [2016 3] [2017 7] [2018 5] [2019 3] [2020 7]]" {
		t.Error("invalid statistic data")
	}
}
