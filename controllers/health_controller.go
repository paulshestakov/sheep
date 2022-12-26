package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	body := make(map[string]string)
	body["message"] = "Status OK"
	c.JSON(http.StatusOK, body)
}