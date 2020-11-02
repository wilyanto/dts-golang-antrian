package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", getSomething)
	router.Run(":8080")
}

func getSomething(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{} {
		"body1":"Get Something Success",
	})
	return
}