package main

import (
	"bookstore-siskastev/config"
	"bookstore-siskastev/controllers"
	"bookstore-siskastev/middleware"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	route := echo.New()
	config.InitDB()
	//user
	route.POST("/user/register", controllers.Register)
	route.POST("/user/login", controllers.Login)

	//book
	books := route.Group("/books", middleware.Auth)
	books.GET("", controllers.GetBooks)
	books.GET("/:id", controllers.GetBookByIsbn)
	books.POST("", controllers.CreateBooks)
	books.PUT("/:id", controllers.UpdateBook)
	books.DELETE("/:id", controllers.DeleteBook)

	route.Start(":8000")
	fmt.Println("server started at localhost:8000")
}

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env")
	}
}
