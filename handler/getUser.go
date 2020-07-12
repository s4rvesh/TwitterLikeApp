package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

//Users data
/*type user struct {
	name  string
	email string
}*/

//GetUser gets the users with specified name
func GetUser(c *gin.Context) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitter")

	if err != nil {
		panic(err.Error())
	}

	name := c.Param("name")
	results, err := db.Query("SELECT name, email FROM users WHERE name=?", name)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user Users

		err = results.Scan(&user.name, &user.email)

		if err != nil {
			panic(err.Error())
		}

		fmt.Print(user.name)
		fmt.Print(user.email)

		c.JSON(200, gin.H{
			"name":  user.name,
			"email": user.email,
		})
	}

	defer results.Close()
	defer db.Close()

}
