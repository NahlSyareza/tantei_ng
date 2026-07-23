package routes

import (
	"tantei-ng/controllers"

	"github.com/gin-gonic/gin"
)

func NgSetRoutes(router *gin.Engine) {
	router.GET("/ng_sets", controllers.GetNgSets)
	router.GET("/ng_set/:oid", controllers.GetNgSet)
	router.POST("/ng_set", controllers.CreateNgSet)
	router.PUT("/ng_set/add/:oid", controllers.AddNgSetItem)
	router.PUT("/ng_set/add_many/:oid", controllers.AddManyNgSetItem)
	router.PUT("/ng_set/remove/:oid", controllers.RemoveNgSetItem)
	router.PUT("/ng_set/remove_many/:oid", controllers.RemoveManyNgSetItem)
}
