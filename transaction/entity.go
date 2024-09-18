package transaction

import (
	"campaign/campaigns"
	"campaign/user"
	"time"
)

type Transaction struct {
	ID         int `gorm:"primaryKey"`
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       int
	PaymentUrl string
	User       user.User           `gorm:"foreignKey:UserId"`
	Campaign   campaigns.Campaigns `gorm:"foreignKey:CampaignId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
