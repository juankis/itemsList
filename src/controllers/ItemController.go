package controllers

import (
	"strconv"

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

func DeleteItem(id string) (string, error) {
	db := db.Connect()
	defer db.Close()
	id_, err := strconv.Atoi(id)
	if err != nil {
		return err.Error(), err
	}
	item := &models.Item{Id: id_}
	err = db.Delete(item)
	if err != nil {
		return err.Error(), err
	}
	return "operacion exitosa", nil
}

func EditItem(id string, title string, description string, picture string) (string, error) {
	db := db.Connect()
	defer db.Close()
	id_, err := strconv.Atoi(id)
	if err != nil {
		return err.Error(), err
	}
	item := models.Item{
		Id:          id_,
		Title:       title,
		Description: description,
	}
	_, err = db.Model(item).Column("title", "description").Update()
	if err != nil {
		return err.Error(), err
	}
	return "operacion exitosa", nil
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
