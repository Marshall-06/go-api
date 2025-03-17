// package main

// import (
// 	"log"
// 	"go-api/database" // Import the database package

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Connect to MySQL database
// 	database.ConnectDB()

// 	// Initialize the Gin router
// 	r := gin.Default()

// 	// Basic route to test
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "Connected to MySQL successfully!"})
// 	})

// 	// Start the server
// 	log.Println("Server is running on http://localhost:8080")
// 	r.Run(":8080")
// }



package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api/database"
	"go-api/routes"
	_ "go-api/docs" // Import Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	database.ConnectDB()

	// Setup Gin router
	r := gin.Default()


	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	// User Routes
	routes.UserRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Server running on port:", port)
	r.Run(":" + port)
}
