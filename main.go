package main

import (
	"fmt"
	"net/http"

	"rest-api-curd-ops/db"
	"rest-api-curd-ops/types"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDynamo()
	if err != nil {
		panic("Unable to connect to DynamoDB")
	}

	r := gin.Default()

	// Apply middleware to all routes
	r.Use(AuthMiddleware())

	r.POST("/item", func(c *gin.Context) {
		var item types.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.CreateItem(item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, item)
		}
	})

	r.GET("/item/:id", func(c *gin.Context) {
		id := c.Param("id")
		item, err := db.GetItem(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusOK, item)
		}
	})

	r.PUT("/item", func(c *gin.Context) {
		var item types.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.UpdateItem(item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, item)
		}
	})

	r.DELETE("/item/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.DeleteItem(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
		}
	})

	fmt.Println("starting server on port 8080")

	r.Run(":8080")
}

// Test Middleware for authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "testToken" { // pass your actual token and validate it against JWK/JWT
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
