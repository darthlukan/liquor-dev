package main

import "github.com/gin-gonic/gin"

func pong(c *gin.Context) {
	c.JSON(200, gin.H{"response": "pong"})
}

func main() {
	srv := gin.Default()
	srv.GET("/ping", pong)
	srv.Run(":8080")
}
