package login

import "errors"

type User struct {
	Username string
	Password string
}

type AuthService struct {
	RegisteredUsers map[string]string
}

func NewAuthService() *AuthService {
	return &AuthService{
		RegisteredUsers: map[string]string{
			"admin": "password123",
			"user1": "123456",
		},
	}
}

func (a *AuthService) Login(username, password string) (*User, error) {
	if storedPassword, exists := a.RegisteredUsers[username]; exists {
		if storedPassword == password {
			return &User{Username: username, Password: password}, nil
		}
		return nil, errors.New("invalid password")
	}
	return nil, errors.New("user not found")
}
