package main

import (
	"TwitterLikeApp/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Starting Twitter like APi")

	r := gin.Default()

	//r.GET("/", handler.Homepage)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"usr": "pass",
	}))

	authorized.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "Gangadhar is  Shaktiman",
		})
	})

	r.GET("/GetUser/:name", handler.GetUser)   //URL Parameter
	r.GET("/TimeLine/:uid", handler.TimeLine)  //URL Parameter
	r.POST("/FollowUser/", handler.FollowUser) //URL Query
	r.POST("/CreateUser", handler.CreateUser)  //JSON POST
	r.DELETE("/DeleteUser/:uid", handler.DeleteUser)
	r.Run()

}
