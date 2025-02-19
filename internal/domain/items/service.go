package items

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

type Service interface {
	CreateItem(ctx context.Context, input CreateReq) error
	GetAll(ctx context.Context) ([]schema.Item, error)
}
