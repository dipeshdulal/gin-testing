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

func doubleRequestHandler(c *gin.Context) {
	req := DoubleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Err: err.Error()})
		return
	}

	if req.Number > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Err: ErrNumberPositive.Error()})
		return
	}

	result := req.Number * 2
	c.JSON(http.StatusOK, DataResponse{Data: result})
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health-check", healthCheckHandler)
	r.POST("/double", doubleRequestHandler)
	return r
}

func main() {
	fmt.Println("running server in :5000")
	getRouter().Run(":5000")
}
