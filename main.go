package main

import (
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/controllers"
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/initializers"
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/migrate"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.Migrate()
}

func main() {

	r := gin.Default()

	r.GET("/", controllers.PostsCreate)

	r.Run()

}
