package auth

import "net/mail"

type User struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Email *mail.Address `json:"email"`
	Thumb string        `json:"thumb"`
}
