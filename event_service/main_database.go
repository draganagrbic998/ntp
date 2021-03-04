package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
)

func openDatabase() {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error = nil
	db, err = gorm.Open("postgres", info)
	if err != nil {
		panic(err)
	}
}

func dropData() {
	db.DropTableIfExists("images")
	db.DropTableIfExists("events")
	db.Exec(createEventsTable)
	db.Exec(createImagesTable)
}

func insertData() {
	sql, err := ioutil.ReadFile("data.sql")
	if err != nil {
		panic(err)
	}
	db.Exec(string(sql))
	db.Exec("ALTER SEQUENCE events_id_seq RESTART WITH 31")
}

func initDatabase() {
	openDatabase()
	defer db.Close()
	dropData()
	insertData()
}
