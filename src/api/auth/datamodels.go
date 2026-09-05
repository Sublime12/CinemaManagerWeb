package auth 

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type MeResponse struct {
	UserID		 uint `json:"user"`
	Username     string `json:"username"`
	Email        string  `json:"email"`
	Name         string `json:"name"`
	IsAdmin      bool    `json:"is_admin"`
}

func ToMeResponse(user User) MeResponse {
	return MeResponse {
		UserID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Name: user.Name,
		IsAdmin: user.IsAdmin,
	}
}
