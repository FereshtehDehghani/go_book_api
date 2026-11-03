package api

import "github.com/gin-gonic/gin"


type Book struct{
	ID int `json:"id" gorm:"primaryKey"`
	Title int `json:"title"`
	Author int `json:"author"`
	Year int `json:"year"`
}

type JsonResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

func ResponseJson(c *gin.Context,status int,message string,data any){

	response:=JsonResponse{
		Status:status,
		Message: message,
		Data: data,
	}

	c.JSON(status,response)

}