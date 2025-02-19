package auth

import (
	"context"
	"fmt"
	"sinarmas-test/internal/domain/auth"
)

func (s *Auth) Login(ctx context.Context, input auth.LoginReq) (string, error) {
	var (
		err error
		jwt string
	)

	user, err := s.data.Get(ctx, input.Username)
	if err != nil {
		return jwt, err
	}

	if user.Password != input.Password {
		return jwt, fmt.Errorf("invalid password")
	}

	//create jwt

	return jwt, err
}
