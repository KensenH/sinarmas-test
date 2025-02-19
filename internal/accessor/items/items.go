package items

import "sinarmas-test/internal/accessor/schema"

type Items struct {
	queries *schema.Queries
}

func New(queries *schema.Queries) Items {
	return Items{
		queries: queries,
	}
}
