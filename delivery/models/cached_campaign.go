package models

// CachedCampaign represents a campaign with its targeting rules
// This is used for caching purposes
type CachedCampaign struct {
	CID     string              `json:"cid"`   // campaign.code
	Img     string              `json:"img"`
	CTA     string              `json:"cta"`
	Include map[string][]string `json:"include"`
	Exclude map[string][]string `json:"exclude"`
}
