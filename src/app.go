package main

import (
	"fmt"
	"log"

	"joyofenergy/src/config"
	"joyofenergy/src/meters"
	priceplans "joyofenergy/src/pricePlans"
	"joyofenergy/src/readings"
	"joyofenergy/src/usage"
)

func main() {
	meterService := meters.NewMeterService()

	// Create some sample meters
	meter1 := meters.NewMeter("smart-meter-0", "electric")
	meter2 := meters.NewMeter("smart-meter-1", "electric")

	meterService.AddMeter(meter1)
	meterService.AddMeter(meter2)

	readingService := readings.NewReadingService(meterService)
	usageService := usage.NewUsageService(readingService)
	pricePlanComparator := priceplans.NewPricePlanComparator(config.PricePlans)

	// Add some sample readings
	err := readingService.StoreReading("smart-meter-0", 10.0)
	if err != nil {
		log.Fatal(err)
	}

	totalUsage, err := usageService.GetTotalUsage("smart-meter-0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Usage: %.2f\n", totalUsage)

	recommendedPlans := pricePlanComparator.GetRecommendedPricePlans(meter1, 2)
	fmt.Println("Recommended Price Plans:", recommendedPlans)
}
