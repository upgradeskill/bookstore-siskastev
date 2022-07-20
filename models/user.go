package models

import (
	"bookstore-siskastev/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func verifyPassword(hashPasswordDb, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasswordDb), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Create() error {
	if err := config.DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) CheckLogin(email, password string) error {
	var err error

	err = config.DB.Where("email = ?", email).Take(&u).Error

	if err != nil {
		return err
	}

	err = verifyPassword(u.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return nil
}
