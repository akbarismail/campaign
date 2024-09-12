package user

import "time"

type User struct {
	ID             int8 `gorm:"primaryKey"`
	Name           string
	Occupation     string
	Email          string
	HashPassword   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
