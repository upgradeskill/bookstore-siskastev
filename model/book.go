package model

import (
	"bookstore-siskastev/config"
)

type Book struct {
	Isbn   string  `gorm:"isbn"`
	Title  string  `gorm:"title"`
	Author string  `gorm:"author"`
	Price  float32 `gorm:"price"`
}

var Books []Book

func All() ([]Book, error) {
	if err := config.DB.Find(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}

func (b *Book) Create() error {
	if err := config.DB.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func GetBookByIsbn(isbn string) ([]Book, error) {
	if err := config.DB.Where("isbn = ?", isbn).Take(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}
