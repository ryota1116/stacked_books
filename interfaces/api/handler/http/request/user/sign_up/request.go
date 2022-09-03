package sign_up

type RequestBody struct {
	Id       int
	UserName string
	Email    string
	Password string
	Avatar   string
	Role     int
}
