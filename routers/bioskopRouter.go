package routers

import (
	"tugas-13/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.AddBioskop)

	return router
}
