package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *KSapplication) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "server health okay",
	})
}
