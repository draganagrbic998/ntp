package main

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getEvents(productID int, page int, size int) ([]event, int) {
	openDatabase()
	defer db.Close()
	var events []event
	var count int

	db.Model(&event{}).
		Where("active = true and product_id = ?", productID).
		Offset(page * size).Limit(size).
		Order("created_on desc").Order("name asc").Find(&events)
	db.Model(&event{}).
		Where("active = true and product_id = ?", productID).
		Count(&count)
	for index, event := range events {
		db.Model(&image{}).Where("event_ref = ?", event.ID).Find(&events[index].Images)
	}

	return events, count
}

func getEvent(id int) (event, int) {
	openDatabase()
	defer db.Close()
	var ev event
	var count int

	db.Model(&event{}).Where("id = ? and active = true", id).Find(&ev).Count(&count)
	db.Model(&image{}).Where("event_ref = ?", ev.ID).Find(&ev.Images)
	return ev, count
}

func createEvent(event event) event {
	openDatabase()
	defer db.Close()

	event.Active = true
	event.CreatedOn = time.Now().UTC().String()
	db.Create(&event)

	var count int
	db.Table("images").Select("max(id)").Row().Scan(&count)

	for index, image := range event.Images {
		count++
		image.EventRef = event.ID
		data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
		path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
		ioutil.WriteFile(path, data, 0644)
		image.Path = serviceURL + "/" + path
		db.Create(&image)
		event.Images[index] = image
	}

	return event
}

func updateEvent(event event) event {
	openDatabase()
	defer db.Close()

	db.Save(&event)
	db.Exec("update images set event_ref = null where event_ref = ?", event.ID)
	var count int
	db.Table("images").Select("max(id)").Row().Scan(&count)

	for index, image := range event.Images {
		count++
		image.EventRef = event.ID
		if image.ID == 0 {
			data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
			path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
			ioutil.WriteFile(path, data, 0644)
			image.Path = serviceURL + "/" + path
			db.Create(&image)
		} else {
			db.Save(&image)
		}
		event.Images[index] = image
	}

	return event
}

func deleteEvent(event event) event {
	openDatabase()
	defer db.Close()

	event.Active = false
	db.Save(&event)
	db.Exec("update images set event_ref = null where event_ref = ?", event.ID)
	return event
}

func statistic(start int, end int) [][2]int {
	openDatabase()
	defer db.Close()
	result := make([][2]int, end-start+1)
	counter := 0

	for i := start; i <= end; i++ {
		var count int
		db.Model(&event{}).Where("substring(created_on, 1, 4) = ?", i).Count(&count)
		result[counter] = [2]int{i, count}
		counter++
	}

	return result
}
