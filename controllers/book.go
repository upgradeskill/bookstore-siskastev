package controllers

import (
	"bookstore-siskastev/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Result struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetBooks(ctx echo.Context) error {
	books, err := models.All()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: books})
}

func CreateBooks(ctx echo.Context) error {
	book := new(models.Book)
	ctx.Bind(&book)
	if err := book.Create(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{
		Message: "Success Create Book",
		Data:    book,
	})
}

func GetBookByIsbn(ctx echo.Context) error {
	id := ctx.Param("id")
	err := models.GetBookByIsbn(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: models.Books})
}

func DeleteBook(ctx echo.Context) error {
	id := ctx.Param("id")
	err := models.GetBookByIsbn(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
	}
	errorDelete := models.Delete(id)
	if errorDelete != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Deleted Data", Data: nil})
}

func UpdateBook(ctx echo.Context) error {
	id := ctx.Param("id")
	errNotFound := models.GetBookByIsbn(id)
	if errNotFound != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: errNotFound.Error(), Data: nil})
	}
	book := new(models.Book)
	book.Isbn = id
	ctx.Bind(&book)
	if err := book.Update(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{
		Message: "Success Update Book",
		Data:    book,
	})

}
