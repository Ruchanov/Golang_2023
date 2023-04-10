package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id          uint   `json:"id" gorm: "primary key"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Cost        int    `json: "cost"`
}
