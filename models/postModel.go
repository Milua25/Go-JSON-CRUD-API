package models

import "github.com/jinzhu/gorm"

type POST struct {
	gorm.Model
	Title string
	Body  string
}
