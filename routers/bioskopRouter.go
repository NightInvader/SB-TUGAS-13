package routers

import (
	"tugas-13/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.AddBioskop)
	router.GET("/bioskop", controllers.GetBioskop)
	router.GET("/bioskop/:id", controllers.GetBioskopByID)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	return router
}
