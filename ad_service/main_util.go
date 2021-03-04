package main

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createToken(userID int) string {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}

func parseToken(request *http.Request) jwt.MapClaims {
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

func imagesNumber() int {
	openDatabase()
	defer db.Close()
	var count int
	db.Table("images").Count(&count)
	return count
}

func detachedImagesNumber() int {
	openDatabase()
	defer db.Close()
	var count int
	db.Model(&image{}).Where("prod_ref is null").Count(&count)
	return count
}

func base64image() string {
	f, _ := os.Open("image.jpeg")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	return "data:image/jpeg;base64, " + encoded
}
