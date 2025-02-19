package auth

import "sinarmas-test/internal/domain/auth"

type Auth struct {
	authSvc auth.Service
}

func New(authSvc auth.Service) Auth {
	return Auth{
		authSvc: authSvc,
	}
}
