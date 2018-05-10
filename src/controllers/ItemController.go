package controllers

import (
	"github.com/juankis/apiItems/src/db"
	"github.com/juankis/apiItems/src/models"
)

func SaveItem(title string, description string, picture string) string {
	db := db.Connect()
	defer db.Close()
	item := &models.Item{Title: title, Description: description, Picture: picture}
	err := db.Insert(item)
	if err != nil {
		return ("Error inserting: " + err.Error())
	} else {
		return "1"
	}
}

func GetItems() []models.Item {
	db := db.Connect()
	defer db.Close()
	var items []models.Item
	err := db.Model(&items).Select()
	if err != nil {
		return items
	}
	return items
}
