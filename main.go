package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		nim := os.Getenv("NIM")
		if nim == "" {
			nim = "NO NIM"
		}
		nama := os.Getenv("NAMA")
		if nama == "" {
			nama = "NO NAMA"
		}
		c.JSON(200, gin.H{
			"NIM":  nim,
			"NAMA": nama,
		})
	})

	router.Run(":8080")
}
