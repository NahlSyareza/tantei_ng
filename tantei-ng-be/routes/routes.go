package routes

import (
	"tantei-ng/controllers"

	"github.com/gin-gonic/gin"
)

func NgSetRoutes(router *gin.Engine) {
	router.GET("/studysets/:owner", controllers.GetNgSets)
	router.GET("/studyset/:studyset", controllers.GetNgSet)
	router.POST("/studyset", controllers.CreateNgSet)
	router.PUT("/studyset/add/:studyset", controllers.AddNgSetItem)
	router.PUT("/studyset/add_many/:studyset", controllers.AddManyNgSetItem)
	router.PUT("/studyset/remove/:studyset", controllers.RemoveNgSetItem)
	router.PUT("/studyset/remove_many/:studyset", controllers.RemoveManyNgSetItem)

	router.POST("/account/register", controllers.RegisterAccount)
	router.POST("/account/login", controllers.LoginAccount)
	// router.POST("/account/owned/:oid", controllers.OwnedSetsByAccount)

	router.POST("/tracker/remove_item", controllers.RemoveItemTracker)
	// param1: oid,
	router.POST("/tracker/reset/:param1", controllers.ResetTracker)
	router.POST("/tracker/refresh/:param1", controllers.RefreshCheckTracker)
}
