package controllers

import (
	"net/http"
	"project_gin/database"
	"project_gin/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateLogin creates a new login
func CreateLogin(c *gin.Context) {
	var login entity.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GetDB().Create(&login).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, login)
}

// GetLogins returns a list of all logins
func GetLogins(c *gin.Context) {
	var logins []entity.Login
	if err := database.GetDB().Find(&logins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logins)
}

// GetLogin returns a single login by ID
func GetLogin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var login entity.Login
	if err := database.GetDB().First(&login, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Login not found"})
		return
	}

	c.JSON(http.StatusOK, login)
}

// UpdateLogin updates a login by ID
func UpdateLogin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var login entity.Login
	if err := database.GetDB().First(&login, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Login not found"})
		return
	}

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GetDB().Save(&login).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, login)
}

// DeleteLogin deletes a login by ID
func DeleteLogin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.GetDB().Where("id = ?", id).Delete(&entity.Login{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
