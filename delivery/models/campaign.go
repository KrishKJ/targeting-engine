package models

// CampaignResponse represents the response structure for a campaign
type Campaign struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Code     string `gorm:"uniqueIndex" json:"code"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	CTA      string `json:"cta"`
	Status   string `json:"status"`
}

// CampaignResponse is the response structure for a campaign
type TargetingRule struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	CampaignID int
	Dimension  string
	Type       string
	Value      string
}
