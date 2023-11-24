package user

type PasswordEncoder interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type BCryptPasswordEncoder struct {
}
