package auth

type Store interface {
	GetUserByCredentials(input AuthenticationInput) (User, error)
}
