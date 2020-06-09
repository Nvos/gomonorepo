package auth

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type AuthenticationInput struct {
	Username string
	Password string
}
