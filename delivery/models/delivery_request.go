package models

// DeliveryRequest represents the request structure for delivery
type DeliveryRequest struct {
	App     string
	Country string
	OS      string
}

// CampaignResponse represents the response structure for a campaign
type CampaignResponse struct {
	CID string `json:"cid"` // use Code, not ID
	Img string `json:"img"`
	CTA string `json:"cta"`
}

