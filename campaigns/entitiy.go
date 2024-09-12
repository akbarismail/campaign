package campaigns

import "time"

type Campaigns struct {
	ID               int
	UserId           int
	Name             string
	Description      string
	ShortDescription string
	CurrentAmount    int
	TotalAmount      int
	Perks            string
	BackerCount      int
	Slug             string
	CampaignImage    []CampaignImages
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImages struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
