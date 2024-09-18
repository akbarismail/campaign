package user

type Formatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageUrl   string `json:"image_url"`
}

func UserFormatter(user User, token string) Formatter {
	formatUser := Formatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImageUrl:   user.AvatarFileName,
	}

	return formatUser
}
