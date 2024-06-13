package database

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/llmos-ai/llmos-controller/pkg/generated/ent"
	"github.com/llmos-ai/llmos-controller/pkg/settings"
)

func GetDBClient() (*ent.Client, error) {
	dbUrl := settings.DatabaseURL.Get()
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	// Create an database.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func RunAutoMigrate(ctx context.Context) (*ent.Client, error) {
	client, err := GetDBClient()
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	// run the auto migration tool.
	if err = client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return client, nil
}
