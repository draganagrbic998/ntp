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

func TestGetEventsEmpty(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/events", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getEventsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var events []event
	json.NewDecoder(rr.Body).Decode(&events)
	if len(events) != 0 {
		t.Error("events list should be empty")
	}

}

func TestGetEventsWithLimit(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/events?size=5&product=1", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getEventsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var events []event
	json.NewDecoder(rr.Body).Decode(&events)
	if len(events) != 1 {
		fmt.Println(len(events))
		t.Error("events list should be length of 1")
	}

	order := events[0].ID == 1 && events[0].Active && events[0].UserID == 1 &&
		events[0].ProductID == 1 && events[0].Name == "Event 1" && events[0].Category == "FAIR" &&
		events[0].Place == "Novi Sad" && events[0].Description == "Description 1" && len(events[0].Images) == 0
	if !order {
		t.Error("invalid events data")
	}
}

func TestGetEventsWithOffset(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/events?page=1&size=5&product=1", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getEventsEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var events []event
	json.NewDecoder(rr.Body).Decode(&events)
	if len(events) != 0 {
		t.Error("events list should be empty")
	}

}

func TestCreateEvent(t *testing.T) {

	req, _ := http.NewRequest("POST", "/api/events", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"ProductID": 1,
		"Name": "test name",
		"Category": "test category",
		"From": "2-2-2020",
		"To": "2-2-2020",
		"Place": "test place",	
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
	handler := http.HandlerFunc(createEventEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ev event
	json.NewDecoder(rr.Body).Decode(&ev)

	order := ev.ID == 31 && ev.Active && ev.UserID == 1 && ev.ProductID == 1 &&
		ev.Name == "test name" && ev.Category == "test category" &&
		ev.From == "2-2-2020" && ev.To == "2-2-2020" && ev.Place == "test place" &&
		ev.Description == "test description" && len(ev.Images) == 2 &&
		ev.Images[0].ID == 1 && ev.Images[0].Path == "http://localhost:8002/image1.jpeg" && ev.Images[0].EventRef == 31 &&
		ev.Images[1].ID == 2 && ev.Images[1].Path == "http://localhost:8002/image2.jpeg" && ev.Images[1].EventRef == 31
	if !order {
		t.Error("invalid event data")
	}

	if count := imagesNumber(); count != 2 {
		t.Error("invalid number of images")
	}

	if count := detachedImagesNumber(); count != 0 {
		t.Error("invalid number of detached images")
	}

	_, err1 := os.Stat("image1.jpeg")
	_, err2 := os.Stat("image2.jpeg")
	if err1 != nil || err2 != nil {
		t.Error("invalid file images")
	}

	os.Remove("image1.jpeg")
	os.Remove("image2.jpeg")
}

func TestCreateEventBadRequest(t *testing.T) {

	req, _ := http.NewRequest("POST", "/api/events", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"ProductID": "",
		"Name": "",
		"Category": "",
		"From": "",
		"To": "",
		"Place": "",	
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
	handler := http.HandlerFunc(createEventEndpoint)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("status code should be 400 Bad Request")
	}
}

func TestUpdateEvent(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/events/31", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"ProductID": 1,
		"Name": "test name 2",
		"Category": "test category",
		"From": "2-2-2020",
		"To": "2-2-2020",
		"Place": "test place",	
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
	r.HandleFunc("/api/events/{id:[0-9]+}", updateEventEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ev event
	json.NewDecoder(rr.Body).Decode(&ev)

	order := ev.ID == 31 && ev.Active && ev.UserID == 1 && ev.ProductID == 1 &&
		ev.Name == "test name 2" && ev.Category == "test category" &&
		ev.From == "2-2-2020" && ev.To == "2-2-2020" && ev.Place == "test place" &&
		ev.Description == "test description" && len(ev.Images) == 2 &&
		ev.Images[0].ID == 3 && ev.Images[0].Path == "http://localhost:8002/image3.jpeg" && ev.Images[0].EventRef == 31 &&
		ev.Images[1].ID == 4 && ev.Images[1].Path == "http://localhost:8002/image4.jpeg" && ev.Images[1].EventRef == 31
	if !order {
		t.Error("invalid event data")
	}

	if count := imagesNumber(); count != 4 {
		t.Error("invalid number of images")
	}

	if count := detachedImagesNumber(); count != 2 {
		t.Error("invalid number of detached images")
	}

	_, err1 := os.Stat("image3.jpeg")
	_, err2 := os.Stat("image4.jpeg")
	if err1 != nil || err2 != nil {
		t.Error("invalid file images")
	}

	os.Remove("image3.jpeg")
	os.Remove("image4.jpeg")
}

func TestUpdateEventBadRequest(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/events/1", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"ProductID": "",
		"Name": "",
		"Category": "",
		"From": "",
		"To": "",
		"Place": "",	
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
	r.HandleFunc("/api/events/{id:[0-9]+}", updateEventEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("status code should be 400 Bad Request")
	}
}

func TestUpdateUpdateNotFound(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/events/40", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"ProductID": 1,
		"Name": "test name",
		"Category": "test category",
		"From": "2-2-2020",
		"To": "2-2-2020",
		"Place": "test place",	
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
	r.HandleFunc("/api/events/{id:[0-9]+}", updateEventEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Error("status code should be 404 Not Found")
	}
}

func TestUpdateEventForbidden(t *testing.T) {

	req, _ := http.NewRequest("PUT", "/api/events/11", bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"ProductID": 1,
		"Name": "test name",
		"Category": "test category",
		"From": "2-2-2020",
		"To": "2-2-2020",
		"Place": "test place",	
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
	r.HandleFunc("/api/events/{id:[0-9]+}", updateEventEndpoint).Methods("PUT")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Error("status code should be 403 Forbidden")
	}
}

func TestDeleteEvent(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/api/events/31", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/events/{id:[0-9]+}", deleteEventEndpoint).Methods("DELETE")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("status code should be 200 OK")
	}

	var ev event
	json.NewDecoder(rr.Body).Decode(&ev)

	order := !ev.Active
	if !order {
		t.Error("invalid event data")
	}

	if count := imagesNumber(); count != 4 {
		t.Error("invalid number of images")
	}

	if count := detachedImagesNumber(); count != 4 {
		t.Error("invalid number of detached images")
	}
}

func TestDeleteEventNotFound(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/api/events/40", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/events/{id:[0-9]+}", deleteEventEndpoint).Methods("DELETE")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Error("status code should be 404 Not Found")
	}
}

func TestDeleteEventForbidden(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/api/events/11", nil)
	token := createToken(1)
	req.Header.Set("Authorization", "JWT "+token)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/events/{id:[0-9]+}", deleteEventEndpoint).Methods("DELETE")
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

	if fmt.Sprintf("%v", result) != "[[2014 4] [2015 2] [2016 4] [2017 2] [2018 9] [2019 7] [2020 2]]" {
		t.Error("invalid statistic data")
	}
}
