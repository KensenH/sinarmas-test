package auth

import "context"

type Service interface {
	Login(ctx context.Context, input LoginReq) (string, error)
}
