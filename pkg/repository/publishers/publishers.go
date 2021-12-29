package publishers

import (
	"github.com/felipear89/go-examples/pkg/model"
	"gorm.io/gorm"
)

type Create func(book *model.Publisher) error

func NewCreate(db *gorm.DB) Create {
	return func(p *model.Publisher) error {
		return db.Create(p).Error
	}
}

type GetAll func() ([]*model.Publisher, error)

func NewGetAll(db *gorm.DB) GetAll {
	return func() ([]*model.Publisher, error) {
		var p []*model.Publisher
		tx := db.Preload("Books").Find(&p)
		if tx.Error != nil {
			return nil, tx.Error
		}
		return p, nil
	}
}
