package main

import (
	"os"
	"tantei-ng/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

// 		// Handle the browser preflight OPTIONS check
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent) // Returns 204 with headers attached
// 			return
// 		}

// 		c.Next()
// 	}
// }

func main() {
	var router *gin.Engine = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Tantei Ng Backend!",
		})
	})

	// routes.WordRoutes(router)
	routes.NgSetRoutes(router)
	// router.POST("/books", createBook)
	// router.GET("/books", getBooks)
	// router.GET("/", getRoot)
	// router.GET("/books/:id", getBookById)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "3000"
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
