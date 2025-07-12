package utils

import (
	"encoding/json"

	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery/models"
)

// ProcessDelivery processes the delivery request and returns matching campaigns
// It fetches active campaigns from the database and checks if they match the targeting rules
func ProcessDelivery(req models.DeliveryRequest) ([]models.CampaignResponse, error) {
	var cached []models.CachedCampaign

	val, err := db.Redis.Get(db.Ctx, "campaigns:active").Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(val), &cached); err != nil {
		return nil, err
	}

	var result []models.CampaignResponse

	for _, c := range cached {
		if matchesFromCache(c, req) {
			result = append(result, models.CampaignResponse{
				CID: c.CID,
				Img: c.Img,
				CTA: c.CTA,
			})
		}
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

// isTargetMatch checks if the campaign matches the targeting rules based on the delivery request
// It returns true if the campaign matches the request parameters
func isTargetMatch(campaignID int, req models.DeliveryRequest) bool {
	var rules []models.TargetingRule
	db.DB.Where("campaign_id = ?", campaignID).Find(&rules)

	include := map[string][]string{}
	exclude := map[string][]string{}

	for _, r := range rules {
		if r.Type == "include" {
			include[r.Dimension] = append(include[r.Dimension], r.Value)
		} else if r.Type == "exclude" {
			exclude[r.Dimension] = append(exclude[r.Dimension], r.Value)
		}
	}

	// Inclusion logic
	if vals, ok := include["country"]; ok && !contains(vals, req.Country) {
		return false
	}
	if vals, ok := include["os"]; ok && !contains(vals, req.OS) {
		return false
	}
	if vals, ok := include["app"]; ok && !contains(vals, req.App) {
		return false
	}

	// Exclusion logic
	if vals, ok := exclude["country"]; ok && contains(vals, req.Country) {
		return false
	}
	if vals, ok := exclude["os"]; ok && contains(vals, req.OS) {
		return false
	}
	if vals, ok := exclude["app"]; ok && contains(vals, req.App) {
		return false
	}

	return true
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
