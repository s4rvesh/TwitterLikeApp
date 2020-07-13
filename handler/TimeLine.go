package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

//Tweets data
type Tweets struct {
	tid  int
	uid  int
	data string
}

//Followers data
type Followers struct {
	uid int
	fid int
}

//TimeLine shows the tweets for the people followed
func TimeLine(c *gin.Context) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitter")

	if err != nil {
		panic(err.Error())
	}

	uid := c.Param("uid")
	results, err := db.Query("SELECT uid,fid FROM followers WHERE uid=?", uid)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var follows Followers
		err = results.Scan(&follows.uid, &follows.fid)

		if err != nil {
			panic(err.Error())
		}

		fmt.Print(follows.uid)
		fmt.Print(follows.fid)

		getTweets, err := db.Query("SELECT data, tid, uid FROM tweets WHERE uid=?", follows.fid)

		if err != nil {
			panic(err.Error())
		}

		for getTweets.Next() {

			var tweet Tweets

			err = getTweets.Scan(&tweet.data, &tweet.tid, &tweet.uid)

			if err != nil {
				panic(err.Error())
			}

			c.JSON(200, gin.H{
				"uid":  tweet.uid,
				"tid":  tweet.tid,
				"data": tweet.data,
			})

		}

	}

	defer results.Close()
	defer db.Close()

}
