package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createHandler(c *gin.Context) {
	reqBody := Link{}
	err := c.Bind(&reqBody)

	if err != nil {
		if err != nil {
			ress := gin.H{
				"error": err.Error(),
			}

			c.JSON(http.StatusBadRequest, ress)
			return
		}
	}

}
