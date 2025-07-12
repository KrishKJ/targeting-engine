package models

type Campaign struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Code     string `gorm:"uniqueIndex" json:"code"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	CTA      string `json:"cta"`
	Status   string `json:"status"`
}

type TargetingRule struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	CampaignID int
	Dimension  string
	Type       string
	Value      string
}
