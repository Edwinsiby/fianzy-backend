package main

import (
	"fianzy/config"
	"fianzy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	app := gin.Default()

	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.MountRoutes(app)
	app.Run(":8080")
}
