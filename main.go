package main

import (
	"restapi-gin-non-database/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", controllers.GetBook)
	router.GET("/books/:Id", controllers.GetBookById)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:Id", controllers.UpdateBook)
	router.DELETE("/books/:Id", controllers.DeleteBook)

	router.Run()
}
