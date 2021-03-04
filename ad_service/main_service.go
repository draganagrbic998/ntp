package main

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getAds(page int, size int, search string) ([]advertisement, int) {
	openDatabase()
	defer db.Close()
	var ads []advertisement
	var count int

	db.Model(&advertisement{}).
		Where("(active = true) and (lower(name) like ? or lower(category) like ? or lower(description) like ?)", search, search, search).
		Offset(page * size).Limit(size).
		Order("created_on desc").Order("name asc").Find(&ads)
	db.Model(&advertisement{}).
		Where("(active = true) and (lower(name) like ? or lower(category) like ? or lower(description) like ?)", search, search, search).
		Count(&count)
	for index, product := range ads {
		db.Model(&image{}).Where("prod_ref = ?", product.ID).Find(&ads[index].Images)
	}

	return ads, count
}

func getMyAds(userID, page int, size int, search string) ([]advertisement, int) {
	openDatabase()
	defer db.Close()
	var ads []advertisement
	var count int

	db.Model(&advertisement{}).
		Where("(user_id = ? and active = true) and (lower(name) like ? or lower(category) like ? or lower(description) like ?)", userID, search, search, search).
		Offset(page * size).Limit(size).
		Order("created_on desc").Find(&ads)
	db.Model(&advertisement{}).
		Where("(user_id = ? and active = true) and (lower(name) like ? or lower(category) like ? or lower(description) like ?)", userID, search, search, search).
		Count(&count)
	for index, product := range ads {
		db.Model(&image{}).Where("prod_ref = ?", product.ID).Find(&ads[index].Images)
	}

	return ads, count
}

func getAd(id int) (advertisement, int) {
	openDatabase()
	defer db.Close()
	var ad advertisement
	var count int

	db.Model(&advertisement{}).Where("id = ? and active = true", id).Find(&ad).Count(&count)
	db.Model(&image{}).Where("prod_ref = ?", ad.ID).Find(&ad.Images)
	return ad, count
}

func createAd(ad advertisement) advertisement {
	openDatabase()
	defer db.Close()

	ad.Active = true
	ad.CreatedOn = time.Now().UTC().String()
	db.Create(&ad)

	var count int
	db.Table("images").Select("max(id)").Row().Scan(&count)

	for i, image := range ad.Images {
		count++
		image.ProdRef = ad.ID
		data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
		path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
		ioutil.WriteFile(path, data, 0644)
		image.Path = serviceURL + "/" + path
		db.Create(&image)
		ad.Images[i] = image
	}

	return ad
}

func updateAd(ad advertisement) advertisement {
	openDatabase()
	defer db.Close()

	db.Save(&ad)
	db.Exec("update images set prod_ref = null where prod_ref = ?", ad.ID)
	var count int
	db.Table("images").Select("max(id)").Row().Scan(&count)

	for i, image := range ad.Images {
		count++
		image.ProdRef = ad.ID
		if image.ID == 0 {
			data, _ := base64.StdEncoding.DecodeString(strings.Split(image.Path, ",")[1])
			path := "image" + strconv.Itoa(count) + "." + strings.Split(strings.Split(image.Path, ";")[0], "/")[1]
			ioutil.WriteFile(path, data, 0644)
			image.Path = serviceURL + "/" + path
			db.Create(&image)
		} else {
			db.Save(&image)
		}
		ad.Images[i] = image
	}

	return ad
}

func deleteAd(ad advertisement) advertisement {
	openDatabase()
	defer db.Close()

	ad.Active = false
	db.Save(&ad)
	db.Exec("update images set prod_ref = null where prod_ref = ?", ad.ID)
	return ad
}

func statistic(start int, end int) [][2]int {
	openDatabase()
	defer db.Close()
	result := make([][2]int, end-start+1)
	counter := 0

	for i := start; i <= end; i++ {
		var count int
		db.Model(&advertisement{}).Where("substring(created_on, 1, 4) = ?", i).Count(&count)
		result[counter] = [2]int{i, count}
		counter++
	}

	return result
}
