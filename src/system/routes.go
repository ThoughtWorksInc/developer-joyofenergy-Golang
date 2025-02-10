package system

import (
	"encoding/json"
	"net/http"

	"joyofenergy/src/controller"
	"joyofenergy/src/repository"
)

func SetupRoutes(
	smartMeterRepo repository.SmartMeterRepositoryInterface,
	pricePlanRepo repository.PricePlanRepositoryInterface,
) *http.ServeMux {
	router := http.NewServeMux()

	// Initialize controllers
	accountService := controller.NewAccountService()
	meterReadingController := controller.NewElectricityReadingController(
		controller.NewMeterReadingManager(smartMeterRepo),
	)
	pricePlanComparatorController := controller.NewPricePlanComparatorController(
		controller.NewPricePlanComparator(smartMeterRepo, pricePlanRepo),
		accountService,
	)

	// System routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to the JoyEnergy",
		})
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("true"))
	})

	// Electricity Reading routes
	router.HandleFunc("/readings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			meterReadingController.StoreReading(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/readings/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			meterReadingController.GetReadings(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Price Plan routes
	router.HandleFunc("/price-plans", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pricePlanComparatorController.GetPricePlans(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/price-plans/compare/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pricePlanComparatorController.ComparePricePlans(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return router
}
