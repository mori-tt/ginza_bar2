package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// CORS middleware
	r.Use(corsMiddleware())

	// Routes
	api := r.Group("/api")
	{
		api.GET("/establishments", getEstablishments)
		api.GET("/establishments/:id", getEstablishment)
		api.POST("/reviews", createReview)
		api.GET("/reviews/:establishmentId", getReviews)
	}

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(authMiddleware())
	{
		admin.POST("/establishments", createEstablishment)
		admin.PUT("/establishments/:id", updateEstablishment)
		admin.DELETE("/establishments/:id", deleteEstablishment)
		admin.GET("/users", getUsers)
		admin.PUT("/users/:id/approve", approveUser)
	}

	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement JWT verification here
		// For now, we'll just check if the Authorization header is present
		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Placeholder handler functions
func getEstablishments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of establishments"})
}

func getEstablishment(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Establishment details", "id": id})
}

func createReview(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Review created"})
}

func getReviews(c *gin.Context) {
	establishmentId := c.Param("establishmentId")
	c.JSON(http.StatusOK, gin.H{"message": "List of reviews", "establishmentId": establishmentId})
}

func createEstablishment(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Establishment created"})
}

func updateEstablishment(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Establishment updated", "id": id})
}

func deleteEstablishment(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Establishment deleted", "id": id})
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func approveUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User approved", "id": id})
}