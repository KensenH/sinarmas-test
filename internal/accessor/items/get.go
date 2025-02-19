package items

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

func (d *Items) GetAll(ctx context.Context) ([]schema.Item, error) {
	return d.queries.GetAllItems(ctx)
}
