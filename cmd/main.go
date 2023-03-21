package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		newUUID, _ := uuid.NewUUID()
		c.JSON(200, gin.H{
			"uuid": newUUID,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
