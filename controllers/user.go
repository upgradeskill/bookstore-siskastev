package controllers

import (
	"bookstore-siskastev/helpers"
	"bookstore-siskastev/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx echo.Context) error {
	user := new(models.User)
	ctx.Bind(&user)
	if errGeneratePass := user.HashPassword(user.Password); errGeneratePass != nil {
		return ctx.JSON(http.StatusInternalServerError, errGeneratePass.Error())
	}
	if err := user.Create(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Register User",
		"data":    user,
	})
}

func Login(ctx echo.Context) error {
	input := new(TokenRequest)
	user := models.User{}
	ctx.Bind(&input)
	err := user.CheckLogin(input.Email, input.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	token, errToken := helpers.GenerateJwtToken(user.Email)
	if errToken != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"data": map[string]interface{}{
			"user":  user,
			"token": token,
		},
	})
}
