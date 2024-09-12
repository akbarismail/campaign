package campaigns

import "time"

type Campaigns struct {
	ID               int `gorm:"primaryKey"`
	UserId           int
	Name             string
	Description      string
	ShortDescription string
	CurrentAmount    int
	TotalAmount      int
	Perks            string
	BackerCount      int
	Slug             string
	CampaignImages   []CampaignImage `gorm:"foreignKey:ID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImage struct {
	ID         int `gorm:"primaryKey"`
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
