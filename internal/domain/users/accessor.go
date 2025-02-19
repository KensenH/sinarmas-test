package users

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

type Accessor interface {
	Get(ctx context.Context, username string) (schema.User, error)
}
