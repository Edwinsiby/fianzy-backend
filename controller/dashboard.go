package controller

import (
	"fianzy/postgres"

	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	data := postgres.GetStats()
	c.JSON(200, data)
}
