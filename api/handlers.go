package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){

	err:= godotenv.Load()

	if err != nil{
		log.Fatal("Failed to connect to database",err)
	}

	dsn := os.Getenv("DB_URL");
	DB,err =gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err != nil{
		log.Fatal("Failed to connect to database",err)
	}

	// migrate the schema
	if err := DB.AutoMigrate(&Book{}); err !=nil{
		log.Fatal("Failed to migrate schema: ",err)
	}

}

func CreateBook(c *gin.Context){

	var book Book

	if  err := c.ShouldBindJSON(&book);err != nil{
		ResponseJson(c,http.StatusBadRequest,"Invalid input",nil)
	}

	DB.Create(&book)
	ResponseJson(c,http.StatusCreated,"Book created successfully",book)

}

func GetBooks(c *gin.Context){
	var books[]Book

	DB.Find(&books)

	ResponseJson(c,http.StatusOK,"Books arrieved successfully",books)
}

func GetBook(c *gin.Context){
	var book Book
	if err := DB.First(&book,c.Param("id")).Error;err !=nil{
		ResponseJson(c,http.StatusNotFound,"Book not found",nil)
		return
	}

	ResponseJson(c,http.StatusOK,"Book retrieved successfully",book)
}

func UpdateBook(c *gin.Context){
	var book Book

	if err :=DB.First(&book,c.Param("id")).Error;err !=nil{
		ResponseJson(c,http.StatusNotFound,"Book not found",nil)
		return
	}

	// Bind request body
	if err :=c.ShouldBindJSON(&book);err != nil{
		ResponseJson(c,http.StatusBadRequest,"Invalid input",nil)
		return
	}

	DB.Save(&book)
	ResponseJson(c,http.StatusOK,"Book updated successfully",book)

}

func DeleteBook(c *gin.Context){
var book Book
if err:=DB.Delete(&book,c.Param("id")).Error;err != nil{
	ResponseJson(c,http.StatusNotFound,"Book not found",nil)
	return
}
ResponseJson(c,http.StatusOK,"Book delete successfully",nil)
}