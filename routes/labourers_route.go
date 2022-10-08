package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/WatipasoChirambo/go-lang-roado/controllers"
)

func LabourerRoute(router *gin.Engine) {
	router.POST("/labourer", controllers.CreateLabourer())
	router.GET("/labourer/:id", controllers.GetLabourer())
	router.PUT("/labourer/:id", controllers.UpdateLabourer())
	router.GET("/labourers", controllers.GetAllLabourers())
}

