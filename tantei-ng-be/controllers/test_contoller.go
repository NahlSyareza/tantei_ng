package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestCtrller(c *gin.Context) {
	rawClaims, exists := c.Get("claims")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Server missing required values!",
		})
		return
	}

	claims := rawClaims.(jwt.MapClaims)

	c.JSON(http.StatusOK, gin.H{
		"msg":     "Success",
		"payload": claims,
	})
}
