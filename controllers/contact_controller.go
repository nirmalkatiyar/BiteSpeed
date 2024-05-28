package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nirmalkatiyar/bitespeed/models"
	"github.com/nirmalkatiyar/bitespeed/services"
	"gorm.io/gorm"
	"net/http"
)

func IdentifyContact(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.IdentifyRequest
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := services.IdentifyContact(db, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"response": response})
	}
}
