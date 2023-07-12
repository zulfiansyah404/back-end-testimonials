package main

import (
	"project/controller"
	"project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/testimonials", controller.Index)
	r.GET("/testimonials/:id", controller.Show)
	r.POST("/testimonials", controller.Create)
	r.PUT("/testimonials/:id", controller.Update)
	r.DELETE("/testimonials/:id", controller.Delete)
	

	r.Run()
}
