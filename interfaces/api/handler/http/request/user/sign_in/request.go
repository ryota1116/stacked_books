package sign_in

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
