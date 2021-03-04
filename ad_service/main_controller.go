package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func getAdsEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	size, _ := strconv.Atoi(request.URL.Query().Get("size"))
	search := "%" + strings.ToLower(request.URL.Query().Get("search")) + "%"
	ads, count := getAds(page, size, search)

	response.Header().Set(enableHeader, firstPageHeader+", "+lastPageHeader)
	response.Header().Set(firstPageHeader, strconv.FormatBool(page == 0))
	response.Header().Set(lastPageHeader, strconv.FormatBool(size*(page+1) >= count))
	json.NewEncoder(response).Encode(ads)
}

func getMyAdsEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	userID := int(claims["user_id"].(float64))
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	size, _ := strconv.Atoi(request.URL.Query().Get("size"))
	search := "%" + strings.ToLower(request.URL.Query().Get("search")) + "%"
	ads, count := getMyAds(userID, page, size, search)

	response.Header().Set(enableHeader, firstPageHeader+", "+lastPageHeader)
	response.Header().Set(firstPageHeader, strconv.FormatBool(page == 0))
	response.Header().Set(lastPageHeader, strconv.FormatBool(size*(page+1) >= count))
	json.NewEncoder(response).Encode(ads)
}

func getAdEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	ad, count := getAd(id)
	if count == 0 {
		response.WriteHeader(404)
		return
	}
	json.NewEncoder(response).Encode(ad)
}

func createAdEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	var ad advertisement
	ad.UserID = int(claims["user_id"].(float64))
	json.NewDecoder(request.Body).Decode(&ad)

	if strings.TrimSpace(ad.Name) == "" || strings.TrimSpace(ad.Category) == "" ||
		strings.TrimSpace(ad.Price) == "" || strings.TrimSpace(ad.Description) == "" {
		response.WriteHeader(400)
		return
	}

	json.NewEncoder(response).Encode(createAd(ad))
}

func updateAdEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	ad, count := getAd(id)
	if count == 0 {
		response.WriteHeader(404)
		return
	}

	if int(claims["user_id"].(float64)) != ad.UserID {
		response.WriteHeader(403)
		return
	}

	ad.Images = nil
	json.NewDecoder(request.Body).Decode(&ad)
	if strings.TrimSpace(ad.Name) == "" || strings.TrimSpace(ad.Category) == "" ||
		strings.TrimSpace(ad.Price) == "" || strings.TrimSpace(ad.Description) == "" {
		response.WriteHeader(400)
		return
	}

	json.NewEncoder(response).Encode(updateAd(ad))
}

func deleteAdEndpoint(response http.ResponseWriter, request *http.Request) {
	claims := parseToken(request)
	if claims == nil {
		response.WriteHeader(401)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	ad, count := getAd(id)
	if count == 0 {
		response.WriteHeader(404)
		return
	}

	if int(claims["user_id"].(float64)) != ad.UserID {
		response.WriteHeader(403)
		return
	}

	json.NewEncoder(response).Encode(deleteAd(ad))
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
