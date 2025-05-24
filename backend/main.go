package main

import (
	"backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	router.Run()
}
