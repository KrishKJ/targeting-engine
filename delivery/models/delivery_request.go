package models

type DeliveryRequest struct {
	App     string
	Country string
	OS      string
}

type CampaignResponse struct {
	CID string `json:"cid"`
	Img string `json:"img"`
	CTA string `json:"cta"`
}
