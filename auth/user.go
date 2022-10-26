package auth

import (
	"net/mail"
	"time"
)

type User struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Email     *mail.Address `json:"email"`
	ValidFrom time.Time     `json:"valid_from"`
	ValidTo   time.Time     `json:"valid_to"`
	Thumb     string        `json:"thumb"`
}
