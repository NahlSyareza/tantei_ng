package main

import (
	"log"
	"os"
	"tantei-ng/db"
	"tantei-ng/jsonwebtoken"
	"tantei-ng/middlewares"
	"tantei-ng/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	jsonwebtoken.GenerateKeyPairs()

	db.DbConnect()

	// EDIT THIS AIGHT
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4173"},
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

	requiresAuthRouter := router.Group("/")
	requiresAuthRouter.Use(middlewares.AuthenticateJWT())
	routes.RequiresAuthRoutes(requiresAuthRouter)

	// routes.WordRoutes(router)
	routes.NgSetRoutes(router)

	err := godotenv.Load()
	if err != nil {
		log.Println("Godotenv failed to load!")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
