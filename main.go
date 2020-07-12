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
	r.POST("/CreateUser", handler.CreateUser)  //JSON POST
	r.GET("/TimeLine/:uid", handler.TimeLine)  //URL Parameter
	r.GET("/GetUser/:name", handler.GetUser)   //URL Parameter
	r.POST("/FollowUser/", handler.FollowUser) //URL Query

	r.Run()

}
