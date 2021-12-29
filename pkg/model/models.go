package model

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name  string `json:"name" validate:"required" gorm:"type:varchar(100);not null;uniqueIndex"`
	Books []Book `json:"books" gorm:"foreignkey:PublisherID"`
}

type Book struct {
	gorm.Model
	Title       string `json:"title" validate:"required" gorm:"uniqueIndex"`
	Author      string `json:"author"`
	PublisherID uint   `json:"publisherId"`
}
