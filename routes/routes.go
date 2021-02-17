package routes

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {

}

// getRoutes{APIVersion}, format can be used make API migration, as in
// to upgrade and support backward compatibility.
// Deprecating API by version
// getRoutesV1: will handle all V1 APIs
// getRoutesV2: will handle all V2 APIs, and so on.
func getRoutesV1() {
	v1 := router.Group("/v1")
}
