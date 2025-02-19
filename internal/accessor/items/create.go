package items

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

func (d *Items) CreateItem(ctx context.Context, args schema.InsertItemParams) (schema.Item, error) {
	return d.queries.InsertItem(ctx, args)
}
