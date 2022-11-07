package auth

type Status int

const (
	Ok Status = iota
	UserNotFound
	PasswordIncorrect
	InternalError
)
