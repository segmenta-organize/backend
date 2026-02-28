package main

import (
	"segmenta/src/configs"
	"segmenta/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load application configuration
	currentConfig := configs.LoadAppConfig()
	if currentConfig == nil {
		panic("[SYSTEM] Failed to load application configuration")
	}

	// Connect to the database
	errorHandler := configs.ConnectDatabase(currentConfig)
	if errorHandler != nil {
		panic("[SYSTEM] Failed to connect to the database : " + errorHandler.Error())
	}

	// Initialize Gin router
	router := gin.Default()
	
	// Setup authentication routes
	routes.SetupAllRoutes(router, currentConfig)

	// Start the server
	if errorHandler := router.Run(":" + currentConfig.BackendPort); errorHandler != nil {
		panic("[SYSTEM] Failed to start the server: " + errorHandler.Error())
	}
}