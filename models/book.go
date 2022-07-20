package models

import (
	"bookstore-siskastev/config"
)

type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
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

func GetBookByIsbn(isbn string) error {
	if err := config.DB.Where("isbn = ?", isbn).Take(&Books).Error; err != nil {
		return err
	}
	return nil
}

func Delete(isbn string) error {
	if err := config.DB.Where("isbn = ?", isbn).Delete(&Books).Error; err != nil {
		return err
	}
	return nil
}

func (b *Book) Update(isbn string) error {
	if err := config.DB.Where("isbn = ?", isbn).Updates(Book{Title: b.Title, Author: b.Author, Price: b.Price}).Error; err != nil {
		return err
	}
	return nil
}
