package models

type Campaign struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	ImageURL string
	CTA      string
	Status   string // ACTIVE or INACTIVE
}

type TargetingRule struct {
	ID         uint   `gorm:"primaryKey"`
	CampaignID string `gorm:"index"`
	Dimension  string // country, os, app
	Type       string // include or exclude
	Value      string
}
