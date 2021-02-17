package routes

import "github.com/gin-gonic/gin"

func inventoryRoutes(rg *gin.RouterGroup) {
	inventoryRG := rg.Group("/inventory")
	inventoryRG.GET("/categories", controllers.InstallOLMS)
	inventoryRG.GET("/category", controllers.InstallOLMS)
	inventoryRG.GET("/category/:id", controllers.InstallOLMS)
	inventoryRG.GET("/category/:id", controllers.InstallOLMS)
	inventoryRG.GET("/category/:id", controllers.InstallOLMS)

}
