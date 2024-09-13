package campaigns

import (
	"campaign/user"
	"time"
)

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
	User             user.User       `gorm:"foreignKey:UserId"`
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
