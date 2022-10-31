package users

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
