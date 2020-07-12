package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FollowUser follows the specified user.
func FollowUser(c *gin.Context) {

	fmt.Println("Endpoint Hit: Follow User")

	uid := c.Query("uid")
	fid := c.Query("fid")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitter")

	if err != nil {
		panic(err.Error())
	}

	_, errQ := db.Query("INSERT INTO followers(uid,fid) VALUES (?,?)", uid, fid)

	if errQ != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"uid": uid,
		"fid": fid,
	})

	defer db.Close()

}
