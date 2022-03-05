package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheckHandler(c *gin.Context) {
	c.Header("x-test-header", "test-value")
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health-check", healthCheckHandler)
	return r
}

func main() {
	fmt.Println("running server in :5000")
	getRouter().Run(":5000")
}
