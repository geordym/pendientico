package model

import "github.com/google/uuid"

type User struct {
	ID    string
	AuthProviderUserID string
	Name  string
	Phone string
	Email string
}

func NewUser(authProviderUserId, name, phone, email string) User {
	return User{
		ID:    uuid.NewString(), 
		AuthProviderUserID: authProviderUserId,
		Name:  name,
		Phone: phone,
		Email: email,
	}
}