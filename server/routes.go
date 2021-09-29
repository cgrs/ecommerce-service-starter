package server

import (
	"github.com/cgrs/ecommerce-service-starter/items"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "OK"})
}

func AddItem(c *gin.Context) {

	var body items.Item
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err, "status": 400})
	}
	i, err := items.AddItem(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err, "status": 500})
	}
	c.JSON(http.StatusCreated, i)
}

func ListItem(c *gin.Context) {
	c.JSON(http.StatusOK, items.ListItems())
}

func AddItemRoutes(rg *gin.RouterGroup) {
	rg.GET("/items", ListItem)
	rg.POST("/items", AddItem)
}