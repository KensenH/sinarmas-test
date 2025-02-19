package users

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

func (d *Users) Get(ctx context.Context, username string) (schema.User, error) {
	return d.queries.GetUser(ctx, username)
}
