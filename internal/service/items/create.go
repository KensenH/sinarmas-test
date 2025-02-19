package items

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
	"sinarmas-test/internal/domain/items"
)

func (s *Items) CreateItem(ctx context.Context, input items.CreateReq) error {
	var err error

	_, err = s.data.CreateItem(ctx, schema.InsertItemParams{
		Name:        input.Name,
		Description: input.Description,
	})

	return err
}
