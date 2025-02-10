package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type PricePlanComparatorController struct {
	pricePlanComparator *PricePlanComparator
	accountService      *AccountService
}

func NewPricePlanComparatorController(
	pricePlanComparator *PricePlanComparator,
	accountService *AccountService,
) *PricePlanComparatorController {
	return &PricePlanComparatorController{
		pricePlanComparator: pricePlanComparator,
		accountService:      accountService,
	}
}

func (c *PricePlanComparatorController) GetPricePlans(w http.ResponseWriter, r *http.Request) {
	pricePlans := c.pricePlanComparator.GetPricePlans()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pricePlans)
}

func (c *PricePlanComparatorController) ComparePricePlans(w http.ResponseWriter, r *http.Request) {
	// Extract smart meter ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid smart meter ID", http.StatusBadRequest)
		return
	}
	smartMeterID := parts[len(parts)-1]

	// Optional limit
	var limit *int
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limitVal, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
		limit = &limitVal
	}

	// Compare price plans
	comparisons, err := c.pricePlanComparator.RecommendCheapestPricePlans(smartMeterID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comparisons)
}
