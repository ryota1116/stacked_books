package sign_up

type RequestBody struct {
	UserName string  `json:"user_name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Avatar   *string `json:"avatar"`
	Role     int     `json:"role"`
}
