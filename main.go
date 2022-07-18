package main

import "github.com/gin-gonic/gin"

type Link struct {
	ID        int    `json:"id"`
	ShortLink string `json:"short_link"`
	LongLink  string `json:"long_link" binding:"required"`
}

func main() {
	CreateDBConnection()

	//importCSV()

	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupRoutes(r *gin.Engine) {

	r.POST("short_link/create", createHandler)
	r.GET("/:id", redirectHandler)

}
