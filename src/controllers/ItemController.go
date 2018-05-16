package controllers

import (
	"strconv"

	"github.com/juankis/itemsList/src/db"
	"github.com/juankis/itemsList/src/models"
)

//SaveItem save item in the data base
func SaveItem(title string, description string, picture string) (string, error) {
	db := db.Connect()
	defer db.Close()
	item := &models.Item{Title: title, Description: description, Picture: picture}
	err := db.Insert(item)
	if err != nil {
		return err.Error(), err
	} else {
		return "operacion exitosa", nil
	}
}

//DeleteItem delete item in the data base
func DeleteItem(id string) (string, error) {
	db := db.Connect()
	defer db.Close()
	idItem, err := strconv.Atoi(id)
	if err != nil {
		return err.Error(), err
	}
	item := &models.Item{Id: idItem}
	err = db.Delete(item)
	if err != nil {
		return err.Error(), err
	}
	return "operacion exitosa", nil
}

//EditItem edit item in the database
func EditItem(id string, title string, description string, picture string) (string, error) {
	db := db.Connect()
	defer db.Close()
	idItem, err := strconv.Atoi(id)
	if err != nil {
		return err.Error(), err
	}
	item := models.Item{
		Id:          idItem,
		Title:       title,
		Description: description,
		Picture:     picture,
	}
	err = db.Update(&item)
	if err != nil {
		return err.Error(), err
	}
	return "operacion exitosa", nil
}

//GetItems this function return list of items in the data base
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
