package items

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
)

func (s *Items) GetAll(ctx context.Context) ([]schema.Item, error) {
	return s.data.GetAll(ctx)
}
