package main

import (
	"go_book_api/api"

	"github.com/gin-gonic/gin"
)


func main(){
	api.InitDB()
	r:=gin.Default()

	// Routes
	r.POST("/book",api.CreateBook)
	r.GET("/books",api.GetBooks)
	r.GET("/book/:id",api.GetBook)
	r.PUT("/book/:id",api.UpdateBook)
	r.DELETE("/book/:id",api.DeleteBook)

	r.Run(":8080")


}