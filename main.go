package main

import (
	"TwitterLikeApp/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/", handler.Homepage)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"uss": "pass",
	}))

	authorized.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
		})
	})

	r.POST("/CreateUser", handler.CreateUser)  //JSON POST
	r.GET("/TimeLine/:uid", handler.TimeLine)  //URL Parameter
	r.GET("/GetUser/:name", handler.GetUser)   //URL Parameter
	r.POST("/FollowUser/", handler.FollowUser) //URL Query

	r.Run()

}
