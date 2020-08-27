package handler

import (
	"TwitterLikeApp/model"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

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

		var follows model.Followers
		err = results.Scan(&follows.Uid, &follows.Fid)

		if err != nil {
			panic(err.Error())
		}

		fmt.Print(follows.Uid)
		fmt.Print(follows.Fid)

		getTweets, err := db.Query("SELECT data, tid, uid FROM tweets WHERE uid=?", follows.Fid)

		if err != nil {
			panic(err.Error())
		}

		for getTweets.Next() {

			var tweet model.Tweets

			err = getTweets.Scan(&tweet.Data, &tweet.Tid, &tweet.Uid)

			if err != nil {
				panic(err.Error())
			}

			c.JSON(200, gin.H{
				"uid":  tweet.Uid,
				"tid":  tweet.Tid,
				"data": tweet.Data,
			})

		}

	}

	defer results.Close()
	defer db.Close()

}
