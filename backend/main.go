package main

import (
	"errors"
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{CreateBatchSize: 1000})
	db := db.Session(&gorm.Session{CreateBatchSize: 1000})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Users{})

	var users = []Users{
		{Firstname: "yaman", Lastname: "verma", Username: "yaman123", Email: "yamanverma123@gmail.com", Password: "yaman123", Confirmpassword: "yaman123", Mobilenumber: "7691092717"},
		{Firstname: "kunal", Lastname: "verma", Username: "kunal123", Email: "kunalverma123@gmail.com", Password: "yaman123", Confirmpassword: "yaman123", Mobilenumber: "9509615887"},
	}
	result := db.Create(users)

	for _, user := range users {
		log.Printf("User ID: %d", user.ID)
	}

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

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Firstname == "" {
		return errors.New("firstname is required")
	}

	if u.Lastname == "" {
		return errors.New("lastname is required")
	}
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}

	if u.Confirmpassword == "" {
		return errors.New("confirm password is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Mobilenumber == "" {
		return errors.New("mobile number is required")
	}

	if u.Password != u.Confirmpassword {
		return errors.New("password and confirm password do not match")
	}

	return nil
}
