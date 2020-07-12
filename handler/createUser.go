package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

//CreateUser adds a new user to system
func CreateUser(c *gin.Context) {

	fmt.Println("Endpoint Hit: Create A new User")

	user := Users{}

	c.Bind(&user)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitter")

	if err != nil {
		panic(err.Error())
	}

	_, errQ := db.Query("INSERT INTO users(uid, name, email, password) VALUES (?,?,?,?)", user.name, user.email, user.password)

	if errQ != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":     user.name,
		"Email":    user.name,
		"Password": user.password,
	})

	defer db.Close()

}
