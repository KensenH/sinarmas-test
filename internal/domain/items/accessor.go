package items

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

type Accessor interface {
	CreateItem(ctx context.Context, args schema.InsertItemParams) (schema.Item, error)
	GetAll(ctx context.Context) ([]schema.Item, error)
}
