package auth

import "sinarmas-test/internal/domain/users"

type Auth struct {
	data users.Accessor
}

func New(data users.Accessor) Auth {
	return Auth{
		data: data,
	}
}

