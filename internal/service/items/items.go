package items

import "sinarmas-test/internal/domain/items"

type Items struct {
	data items.Accessor
}

func New(data items.Accessor) Items {
	return Items{
		data: data,
	}
}
