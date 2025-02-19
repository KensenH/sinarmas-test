package boot

import (
	"context"
	"sinarmas-test/internal/accessor/schema"
	"sinarmas-test/internal/accessor/users"
	"sinarmas-test/internal/server"
	"sinarmas-test/internal/service/auth"

	itemsData "sinarmas-test/internal/accessor/items"
	authHandler "sinarmas-test/internal/server/api/http/auth"
	itemsHandler "sinarmas-test/internal/server/api/http/items"
	itemsService "sinarmas-test/internal/service/items"

	"github.com/jackc/pgx/v5"
)

func Start(ctx context.Context) error {
	var err error

	conn, err := pgx.Connect(ctx, "postgres://postgres:mysecretpassword@localhost:5432/postgres")
	if err != nil {
		return err
	}

	queries := schema.New(conn)

	userD := users.New(queries)
	authS := auth.New(&userD)
	authH := authHandler.New(&authS)

	itemD := itemsData.New(queries)
	itemS := itemsService.New(&itemD)
	itemH := itemsHandler.New(&itemS)

	s := server.New(&authH, &itemH)

	return s.Start()
}
