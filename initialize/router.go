package initialize

import (
	"github.com/gin-gonic/gin"
	routerV1 "ms-inventory/router/v1"
)

// router function
func Routers() *gin.Engine {
	var router = gin.Default()
	initV1(router)
	return router
}

// Routers initialization
// Multiple API versions can be initialized
func initV1(router *gin.Engine) {
	v1Group := router.Group("v1")
	{
		routerV1.InitInventoryRouter(v1Group)
	}
}
