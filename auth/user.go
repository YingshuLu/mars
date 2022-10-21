package auth

import "net/mail"

type User struct {
	ID    string
	Name  string
	Email *mail.Address
	Thumb []byte
}
