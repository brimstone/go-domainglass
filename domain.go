package main

import "github.com/gin-gonic/gin"

// Domain Handle main route
func apiDomain(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
