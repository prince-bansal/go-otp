package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	appConfig := InitDependencies()

	appConfig.Router.InitRoutes(router)
	err := router.Run()
	if err != nil {
		_ = fmt.Errorf("getting error while starting server")
		return
	}
}
