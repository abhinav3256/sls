package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func redirectHandler(c *gin.Context) {

	id, _ := c.Params.Get("id")

	fmt.Println(id)
	idd, _ := strconv.Atoi(id)
	long_url, expire_at := getLongURL(idd)
	var currentTime = time.Now().UTC()
	expire_time, _ := time.Parse(time.RFC3339, expire_at)
	if currentTime.After(expire_time) {
		res := gin.H{
			"success": false,
			"message": "Page Expire",
		}
		c.JSON(http.StatusNotFound, res)
		c.Abort()
		return
	}

	c.Redirect(http.StatusMovedPermanently, long_url)

}

func getLongURL(id int) (string, string) {

	long_url := ""
	expire_at := ""
	SQL := `SELECT "expire_at","long_link" FROM "link" WHERE id=$1`

	rows := DB.QueryRow(SQL, id)

	rows.Scan(&expire_at, &long_url)
	fmt.Println(long_url)

	return long_url, expire_at
}
