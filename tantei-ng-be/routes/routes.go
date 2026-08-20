package routes

import (
	"tantei-ng/controllers"

	"github.com/gin-gonic/gin"
)

func NgSetRoutes(router *gin.Engine) {
	router.GET("/ng_sets/:owner", controllers.GetNgSets)
	router.GET("/ng_set/:studyset", controllers.GetNgSet)
	router.POST("/ng_set", controllers.CreateNgSet)
	router.PUT("/ng_set/add/:studyset", controllers.AddNgSetItem)
	router.PUT("/ng_set/add_many/:studyset", controllers.AddManyNgSetItem)
	router.PUT("/ng_set/remove/:studyset", controllers.RemoveNgSetItem)
	router.PUT("/ng_set/remove_many/:studyset", controllers.RemoveManyNgSetItem)

	router.POST("/account/register", controllers.RegisterAccount)
	router.POST("/account/login", controllers.LoginAccount)
	// router.POST("/account/owned/:oid", controllers.OwnedSetsByAccount)

	router.POST("/tracker/remove_item", controllers.RemoveItemTracker)
	// param1: oid,
	router.POST("/tracker/reset/:param1", controllers.ResetTracker)
	router.POST("/tracker/refresh/:param1", controllers.RefreshCheckTracker)
}
