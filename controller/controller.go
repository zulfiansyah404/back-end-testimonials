package controller

import (

	// "encoding/json"
	"net/http"

	"project/models"

	// "gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var testimonials []models.Testimonials

	models.DB.Find(&testimonials)

	c.JSON(http.StatusOK, testimonials)
}

func Show(c *gin.Context) {
	var testimonials models.Testimonials

	if err := models.DB.Where("id = ?", c.Param("id")).First(&testimonials).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Testimonials not found!"})
		return
	}

	c.JSON(http.StatusOK, testimonials)
}

func Create(c *gin.Context) {
	var input models.Testimonials

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testimonials := models.Testimonials{Image: input.Image, Text: input.Text, Authors: input.Authors, Rating: input.Rating, IsCompany: input.IsCompany}
	if err := models.DB.Create(&testimonials).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testimonials)
}

func Update(c *gin.Context) {
	var testimonials models.Testimonials

	if err := models.DB.Where("id = ?", c.Param("id")).First(&testimonials).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Testimonials not found!"})
		return
	}

	var input models.Testimonials

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&testimonials).Updates(input)

	c.JSON(http.StatusOK, testimonials)
}

func Delete(c *gin.Context) {
	var testimonials models.Testimonials

	if err := models.DB.Where("id = ?", c.Param("id")).First(&testimonials).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Testimonials not found!"})
		return
	}

	models.DB.Delete(&testimonials)

	c.JSON(http.StatusOK, testimonials)
}
