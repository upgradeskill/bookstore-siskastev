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
	route.GET("/books/:id", getBookByIsbn)
	route.POST("/books", createBooks)
	route.PUT("/books/:id", updateBook)
	route.DELETE("/books/:id", deleteBook)
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
	err := model.GetBookByIsbn(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: model.Books})
}

func deleteBook(ctx echo.Context) error {
	id := ctx.Param("id")
	err := model.GetBookByIsbn(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
	}
	errorDelete := model.Delete(id)
	if errorDelete != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Deleted Data", Data: nil})
}

func updateBook(ctx echo.Context) error {
	id := ctx.Param("id")
	errNotFound := model.GetBookByIsbn(id)
	if errNotFound != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: errNotFound.Error(), Data: nil})
	}
	book := new(model.Book)
	book.Isbn = id
	ctx.Bind(&book)
	if err := book.Update(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{
		Message: "Success Update Book",
		Data:    book,
	})

}
