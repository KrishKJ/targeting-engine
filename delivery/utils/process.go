package utils

import (
	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery/models"
)

func ProcessDelivery(req models.DeliveryRequest) ([]models.CampaignResponse, error) {
	var campaigns []models.Campaign
	db.DB.Where("status = ?", "ACTIVE").Find(&campaigns)

	var results []models.CampaignResponse

	for _, camp := range campaigns {
		if isTargetMatch(camp.ID, req) {
			results = append(results, models.CampaignResponse{
				CID: camp.ID,
				Img: camp.ImageURL,
				CTA: camp.CTA,
			})
		}
	}

	return results, nil
}

func isTargetMatch(campaignID string, req models.DeliveryRequest) bool {
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

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

