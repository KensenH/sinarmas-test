package users

import "sinarmas-test/internal/accessor/schema"

type Users struct {
	queries *schema.Queries
}

func New(queries *schema.Queries) Users {
	return Users{
		queries: queries,
	}
}
