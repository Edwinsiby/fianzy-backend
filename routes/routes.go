package routes

import (
	"fianzy/controller"

	"github.com/gin-gonic/gin"
)

func MountRoutes(app *gin.Engine) {
	app.GET("stats", controller.GetStats)
	app.GET("transactions", controller.GetTransactions)
	app.POST("link", controller.LinkFunds)

	app.POST("add/transactions", controller.CreateTransactions)
	app.POST("repay/:borrow_id", controller.RepayBorrow)
}
