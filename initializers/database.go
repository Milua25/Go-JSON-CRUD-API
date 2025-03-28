package initializers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB_URL := os.Getenv("DB_URL")
	DB, err = gorm.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	//defer DB.Close()
}
