package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Homepage of the application
func Homepage(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello!!!!",
	})

	fmt.Println("Endpoint Hit: HomePage")

}
