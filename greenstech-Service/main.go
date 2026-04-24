package main

import (
	"fmt"
	"greenstech/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found (continuing...)")
	}

	// Create Gin router
	r := gin.Default()

	// Trust proxies (optional)
	r.SetTrustedProxies(nil)

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize your routes
	routes.InitLoginRoutes(r)
	routes.InitProfileRoutes(r)
	routes.InitRegistationRoutes(r)
	routes.InitFileHandlerRoutes(r)
	routes.InitSubtrainerRoutes(r)
	routes.InitGroupsRoutes(r)
	routes.InitStudentRoutes(r)
	routes.InitFilterRoutes(r)
	routes.InitSyllabusRoutes(r)
	routes.InitSessionRoutes(r)
	routes.InitPermissionRoutes(r)
	routes.InitLeaveRoutes(r)
	routes.InitPermissionRequestRoutes(r)

	// Health check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// ✅ PORT handling (correct placement)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("✅ Server is Running at Port:", port)

	// Run server
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
