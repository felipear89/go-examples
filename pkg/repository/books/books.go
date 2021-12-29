package books

import (
	"github.com/felipear89/go-examples/pkg/model"
	"gorm.io/gorm"
)

type Create func(book *model.Book) error

func NewCreate(db *gorm.DB) Create {
	return func(book *model.Book) error {
		return db.Create(&book).Error
	}
}

type GetAll func() ([]*model.Book, error)

func NewGetAll(db *gorm.DB) GetAll {
	return func() ([]*model.Book, error) {
		var books []*model.Book
		tx := db.Find(&books)
		if tx.Error != nil {
			return nil, tx.Error
		}
		return books, nil
	}
}
