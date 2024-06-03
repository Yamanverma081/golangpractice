package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	ID              int
	Firstname       string
	Lastname        string
	Username        string
	Password        string
	Confirmpassword string
	Email           string
	Mobilenumber    string
}

var db *gorm.DB

func main() {

	dsn := "root:yaman@tcp(localhost:3333)/daily_diary?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Users{})

	users := []*Users{
		{Firstname: "yaman", Lastname: "verma"},
		{Firstname: "kunal", Lastname: "verma"},
	}
	result := db.Create(users)

	result.RowsAffected = db.RowsAffected

	router := gin.Default()
	router.GET("/users", selectedUsers)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func selectedUsers(users *gin.Context) {
	var selectedUsers []Users
	db.Select("firstname", "lastname", "username", "email", "mobilenumber").Find(&selectedUsers)
	users.JSON(http.StatusOK, selectedUsers)
}
