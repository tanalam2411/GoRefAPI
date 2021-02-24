package v1

import (
	"github.com/gin-gonic/gin"
	v1 "ms-inventory/pkg/controllers/v1/inventory"
)

func InitInventoryRouter(rg *gin.RouterGroup) {
	inventoryRG := rg.Group("/inventory")
	{
		inventoryRG.POST("/categories", v1.CreateInventoryCategory)
		inventoryRG.GET("/category/:id", v1.GetInventoryCategory)
		inventoryRG.PUT("/category/:id", v1.UpdateInventoryCategory)
		inventoryRG.DELETE("/category/:id", v1.DeleteInventoryCategory)
	}
}
