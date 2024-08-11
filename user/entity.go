package user

import "time"

type User struct {
	ID             int8      `json:"id"`
	Name           string    `json:"name"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	HashPassword   string    `json:"hash_password"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
