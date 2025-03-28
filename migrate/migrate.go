package migrate

import (
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/initializers"
	"github.com/Golang-Personal-Projects/Go-Projects/05-Go-JSON-CRUD-API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Migrate() {
	initializers.DB.AutoMigrate(&models.POST{})
}
