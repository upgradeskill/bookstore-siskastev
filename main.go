package main

import (
	"bookstore-siskastev/config"
	"bookstore-siskastev/model"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Result struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	route := echo.New()
	config.InitDB()
	route.GET("/books", getBooks)
	route.POST("/books", createBooks)
	route.GET("/book/:id", getBookByIsbn)
	route.Start(":8000")
	fmt.Println("server started at localhost:8000")

}

func getBooks(ctx echo.Context) error {
	books, err := model.All()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: books})
}

func createBooks(ctx echo.Context) error {
	book := new(model.Book)
	ctx.Bind(&book)
	if err := book.Create(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{
		Message: "Success Create Book",
		Data:    book,
	})
}

func getBookByIsbn(ctx echo.Context) error {
	id := ctx.Param("id")
	book, err := model.GetBookByIsbn(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, err)
		}

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: book})
}
