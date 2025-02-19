package items

import (
	"sinarmas-test/internal/domain/auth"
	"sinarmas-test/internal/domain/items"
)

type Items struct {
	authSvc auth.Service
	itemSvc items.Service
}

func New(itemSvc items.Service) Items {
	return Items{
		itemSvc: itemSvc,
	}
}
