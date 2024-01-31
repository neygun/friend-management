package authentication

type Service interface {
	IsBlacklisted(token string) (bool, error)
}

// type service struct {
// 	cacheRepo authentication.Repository
// }

// func New(cacheRepo authentication.Repository) Service {
// 	return service{
// 		cacheRepo: cacheRepo,
// 	}
// }
