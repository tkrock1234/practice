package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tkrock1234/practice/models"
)

// main is func
func main() {
	s1 := models.HelloWorld("hoge")
	fmt.Println(s1)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message", "pong",
		})
	})
	r.Run()
}
