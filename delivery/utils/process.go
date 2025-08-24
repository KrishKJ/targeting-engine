package utils

import (
	"encoding/json"
	"runtime"
	"sync"

	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery/models"
)

// ProcessDelivery processes the delivery request and returns matching campaigns
// It fetches active campaigns from the database and checks if they match the targeting rules
func ProcessDelivery(req models.DeliveryRequest) ([]models.CampaignResponse, error) {
	val, err := db.Redis.Get(db.Ctx, "campaigns:active").Result()
	if err != nil {
		return nil, err
	}

	var campaigns []models.CachedCampaign
	if err := json.Unmarshal([]byte(val), &campaigns); err != nil {
		return nil, err
	}

	// Create input & output channels
	numWorkers := runtime.NumCPU() * 2 // tweak for benchmarking
	in := make(chan models.CachedCampaign, len(campaigns))
	out := make(chan models.CampaignResponse, len(campaigns))

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for campaign := range in {
				if matchesFromCache(campaign, req) {
					out <- models.CampaignResponse{
						CID: campaign.CID,
						Img: campaign.Img,
						CTA: campaign.CTA,
					}
				}
			}
		}()
	}

	// Fan-out: send work
	go func() {
		for _, c := range campaigns {
			in <- c
		}
		close(in)
	}()

	// Fan-in: collect results
	go func() {
		wg.Wait()
		close(out)
	}()

	var result []models.CampaignResponse
	for match := range out {
		result = append(result, match)
	}

	return result, nil
}

// matchesFromCache checks if the cached campaign matches the delivery request
// It verifies inclusion and exclusion rules based on the request parameters
// Returns true if the campaign matches the request, false otherwise
func matchesFromCache(c models.CachedCampaign, req models.DeliveryRequest) bool {
	// Inclusion
	for dim, vals := range c.Include {
		if !contains(vals, getDimValue(dim, req)) {
			return false
		}
	}

	// Exclusion
	for dim, vals := range c.Exclude {
		if contains(vals, getDimValue(dim, req)) {
			return false
		}
	}

	return true
}

// getDimValue retrieves the value of a specific dimension from the delivery request
// It returns the corresponding value based on the dimension type (country, os, app)
func getDimValue(dim string, req models.DeliveryRequest) string {
	switch dim {
	case "country":
		return req.Country
	case "os":
		return req.OS
	case "app":
		return req.App
	default:
		return ""
	}
}

// contains checks if a slice contains a specific value
// It is used to verify if the request parameters match the targeting rules
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
