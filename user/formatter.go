package user

type Formatter struct {
	ID         int8   `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func UserFormatter(user User, token string) Formatter {
	formatUser := Formatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatUser
}
