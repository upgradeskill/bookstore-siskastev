package main

import (
	"bookstore-siskastev/config"
	"bookstore-siskastev/controllers"
	"bookstore-siskastev/middleware"
	"bookstore-siskastev/models"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
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
	//user
	route.POST("/user/register", controllers.Register)
	route.POST("/user/login", controllers.Login)
	//book

	route.GET("/books", middleware.Auth(getBooks))
	route.GET("/books/:id", middleware.Auth(getBookByIsbn))
	route.POST("/books", middleware.Auth(createBooks))
	route.PUT("/books/:id", middleware.Auth(updateBook))
	route.DELETE("/books/:id", middleware.Auth(deleteBook))

	route.Start(":8000")
	fmt.Println("server started at localhost:8000")

}

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env")
	}
}

func getBooks(ctx echo.Context) error {
	books, err := models.All()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: books})
}

func createBooks(ctx echo.Context) error {
	book := new(models.Book)
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
	err := models.GetBookByIsbn(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: models.Books})
}

func deleteBook(ctx echo.Context) error {
	id := ctx.Param("id")
	err := models.GetBookByIsbn(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
	}
	errorDelete := models.Delete(id)
	if errorDelete != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Deleted Data", Data: nil})
}

func updateBook(ctx echo.Context) error {
	id := ctx.Param("id")
	errNotFound := models.GetBookByIsbn(id)
	if errNotFound != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: errNotFound.Error(), Data: nil})
	}
	book := new(models.Book)
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
