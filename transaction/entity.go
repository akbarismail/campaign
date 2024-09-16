package transaction

import (
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
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
