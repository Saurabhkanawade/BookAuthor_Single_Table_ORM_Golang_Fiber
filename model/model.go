package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

type Author struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
}
