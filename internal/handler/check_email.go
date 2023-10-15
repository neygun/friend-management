package handler

import "net/mail"

// IsEmail checks if a string is an email
func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
