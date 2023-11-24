package user

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password
func (b BCryptPasswordEncoder) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with its possible plaintext equivalent
func (b BCryptPasswordEncoder) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
