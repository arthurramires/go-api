package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize router
	router := gin.Default()

	//Initialize routes
	initializeRoutes(router)

	//Start de server
	router.Run(":8086")
}
