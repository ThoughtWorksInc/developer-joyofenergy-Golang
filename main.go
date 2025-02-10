package main

import (
	"log"
	"net/http"

	"joyofenergy/src/repository"
	"joyofenergy/src/system"
)

func main() {
	// Initialize repositories
	smartMeterRepo := repository.NewSmartMeterRepository()
	pricePlanRepo := repository.NewPricePlanRepository()

	// Setup routes
	router := system.SetupRoutes(smartMeterRepo, pricePlanRepo)

	// Start server
	log.Println("Server starting on :8020")
	if err := http.ListenAndServe(":8020", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
