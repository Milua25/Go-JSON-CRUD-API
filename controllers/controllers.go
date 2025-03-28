package controllers

import (
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/initializers"
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.POST{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//
	c.JSON(200, gin.H{
		"post": post,
	})
}
