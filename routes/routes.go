package routes

import (
	"fianzy/controller"

	"github.com/gin-gonic/gin"
)

func MountRoutes(app *gin.Engine) {
	app.GET("stats", controller.GetStats)
	app.POST("add/credit", controller.CreateCredit)
	app.POST("add/debit", controller.CreateDebit)
}
