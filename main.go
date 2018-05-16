package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juankis/itemsList/src/controllers"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/form_post", func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
		title := c.PostForm("title")
		description := c.PostForm("description")
		picture := moverArchivo(c)
		msj, err := controllers.SaveItem(title, description, picture)
		fmt.Printf("title: %s;", title)
		c.JSON(200, gin.H{
			"status": msj,
			"error":  err,
		})
	})
	router.POST("/edit_item", handleEditItem)
	router.GET("/get_items", func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"items": controllers.GetItems(),
		})
	})

	router.GET("/delete_item", func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		msj, err := controllers.DeleteItem(c.Query("id"))
		c.JSON(200, gin.H{
			"response": msj,
			"error":    err,
		})
	})

	router.Static("/public/js", "./public/js")
	router.Static("/public/styles", "./public/styles")
	router.Static("/pictures", "./pictures")
	router.LoadHTMLGlob("public/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":9000")
}

func moverArchivo(c *gin.Context) string {
	path := "./pictures/"
	file, err := c.FormFile("picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	}
	filename := strings.Replace(file.Filename, " ", "", -1)
	nameFile := strconv.Itoa(rand.Intn(10000)) + "_" + filename
	if err := c.SaveUploadedFile(file, path+nameFile); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	c.String(http.StatusOK, fmt.Sprintf("El archivo %s ha sido trasladado con exito", filename))
	return nameFile
}

func handleEditItem(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Next()
	id := c.PostForm("id")
	title := c.PostForm("title")
	description := c.PostForm("description")
	picture := moverArchivo(c)
	msj, err := controllers.EditItem(id, title, description, picture)

	c.JSON(200, gin.H{
		"status": "posted",
		"title":  title,
		"msj":    msj,
		"err":    err,
	})
}
