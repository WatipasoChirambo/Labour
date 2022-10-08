package main

import (
	"github.com/gin-gonic/gin"
	"github.com/WatipasoChirambo/go-lang-roado/configs"
	"github.com/WatipasoChirambo/go-lang-roado/routes"
)


func main()  {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.LabourerRoute(router)

	router.Run("localhost:8000")
}
