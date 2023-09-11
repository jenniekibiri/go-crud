package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jenniekibiri/crud-golang/controllers"
	
)


func PersonRouter (router *gin.Engine){
	routes:=router.Group("/api/v1/persons")
	routes.POST("/", controllers.PersonCreate)
	routes.GET("/", controllers.PersonGet)
	routes.PUT("/:id", controllers.PersonUpdate)
	routes.GET("/:id", controllers.PersonGetById)
	routes.DELETE("/:id", controllers.PersonDelete)
}

