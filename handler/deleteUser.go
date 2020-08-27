package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

//DeleteUser delets a  user
func DeleteUser(c *gin.Context) {

	fmt.Println("Endpoint Hit: Delete a User")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitter")

	if err != nil {
		panic(err.Error())
	}

	uid := c.Param("uid")
	_, err = db.Query("DELETE FROM users WHERE uid=?", uid)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"uid":     uid,
		"message": "user deleted",
	})

	defer db.Close()

}
