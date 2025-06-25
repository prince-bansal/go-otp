package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func initServer() {
	router := gin.Default()
	err := router.Run()
	if err != nil {
		_ = fmt.Errorf("getting error while starting server")
		return
	}
}
