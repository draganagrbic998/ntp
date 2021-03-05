package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func getEventsEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	size, _ := strconv.Atoi(request.URL.Query().Get("size"))
	productID, _ := strconv.Atoi(request.URL.Query().Get("product"))
	events, count := getEvents(productID, page, size)

	response.Header().Set(enableHeader, firstPageHeader+", "+lastPageHeader)
	response.Header().Set(firstPageHeader, strconv.FormatBool(page == 0))
	response.Header().Set(lastPageHeader, strconv.FormatBool(size*(page+1) >= count))
	json.NewEncoder(response).Encode(events)
}

func createEventEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	var event event
	event.UserID = int(claims["user_id"].(float64))
	json.NewDecoder(request.Body).Decode(&event)

	if event.ProductID == 0 ||
		strings.TrimSpace(event.Name) == "" || strings.TrimSpace(event.Category) == "" || strings.TrimSpace(event.From) == "" ||
		strings.TrimSpace(event.To) == "" || strings.TrimSpace(event.Place) == "" || strings.TrimSpace(event.Description) == "" {
		response.WriteHeader(400)
		return
	}

	json.NewEncoder(response).Encode(createEvent(event))
}

func updateEventEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	event, count := getEvent(id)
	if count == 0 {
		response.WriteHeader(404)
		return
	}

	if int(claims["user_id"].(float64)) != event.UserID {
		response.WriteHeader(403)
		return
	}

	event.Images = nil
	json.NewDecoder(request.Body).Decode(&event)

	if strings.TrimSpace(event.Name) == "" || strings.TrimSpace(event.Category) == "" || strings.TrimSpace(event.From) == "" ||
		strings.TrimSpace(event.To) == "" || strings.TrimSpace(event.Place) == "" || strings.TrimSpace(event.Description) == "" {
		response.WriteHeader(400)
		return
	}

	json.NewEncoder(response).Encode(updateEvent(event))
}

func deleteEventEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	event, count := getEvent(id)
	if count == 0 {
		response.WriteHeader(404)
		return
	}

	if int(claims["user_id"].(float64)) != event.UserID {
		response.WriteHeader(403)
		return
	}

	json.NewEncoder(response).Encode(deleteEvent(event))
}

func statisticEndpoint(response http.ResponseWriter, request *http.Request) {
	start, _ := strconv.Atoi(mux.Vars(request)["start"])
	end, _ := strconv.Atoi(mux.Vars(request)["end"])

	if start >= end {
		response.WriteHeader(400)
		return
	}

	json.NewEncoder(response).Encode(statistic(start, end))
}
