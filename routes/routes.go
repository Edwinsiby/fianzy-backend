package routes

import (
	"fianzy/controller"

	"github.com/gin-gonic/gin"
)

func MountRoutes(app *gin.Engine) {
	app.GET("stats", controller.GetStats)
	app.POST("add/transactions", controller.CreateTransactions)
	app.GET("get/transactions", controller.GetTransactions)
	app.POST("link", controller.LinkFunds)

	app.POST("repay/:id", controller.Setlement)
}
