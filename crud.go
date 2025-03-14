package main

import (
	"log"
	"net/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID uint `gorm: "primary_key" json: "id"`
	Name string `json: "name"`
	Price float64 `json: "price"`
}

var db *gorm.DB


func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func getProduct(c *gin.Context) {
	id := c.param("id")
	var product Product
	productData := db.First(&product, id)
	if productData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func createProduct(c *gin.Context) {
	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	productData := db.First(&product, id)
	if productData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&product).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	productData := db.First(&product, id)
	if productData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

