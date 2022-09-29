package main

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/controller"
)

func main() {
	router := gin.Default()
	controller.ConfigureLayers(router)
	router.Run(":8000")
}
