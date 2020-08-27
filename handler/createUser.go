package handler

import (
	"TwitterLikeApp/model"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser adds a new user to system
func CreateUser(c *gin.Context) {

	fmt.Println("Endpoint Hit: Create A new User")

	user := model.Users{}

	err := c.Bind(&user)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(user.Name)

	fmt.Println(user.Email)
	fmt.Println(user.Password)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitter")

	if err != nil {
		panic(err.Error())
	}

	_, errQ := db.Query("INSERT INTO users(name, email, password) VALUES (?,?,?)", user.Name, user.Email, user.Password)

	if errQ != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":     user.Name,
		"Email":    user.Email,
		"Password": user.Password,
	})

	defer db.Close()

}
