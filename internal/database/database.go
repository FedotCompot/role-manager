package database

import (
	"context"
	"log"
	"log/slog"
	"role-manager-bot/internal/config"
	"role-manager-bot/internal/database/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
)

var db *bun.DB

func Connect(ctx context.Context) error {

	log.Println(config.Data.DatabaseConnection)
	config, err := pgxpool.ParseConfig(config.Data.DatabaseConnection)
	if err != nil {
		return err
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}

	sqldb := stdlib.OpenDBFromPool(pool)
	db = bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	migrator := migrate.NewMigrator(db, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return err
	}
	if _, err := migrator.Migrate(ctx); err != nil {
		slog.Error("Failed to apply migrations", "error", err)
		return err
	}

	return db.Ping()
}

func Close() error {
	slog.Debug("Closing database")
	if db != nil {
		return db.Close()
	}
	return nil
}

func InsertAndReturnID(ctx context.Context, model any) (returningID int64, err error) {
	err = db.NewInsert().Returning("id").Model(model).ExcludeColumn("id").Scan(ctx, &returningID)
	return
}

func InsertAndReturnAll[T any](ctx context.Context, model *T) (*T, error) {
	_, err := db.NewInsert().Returning("*").Model(model).ExcludeColumn("id").Exec(ctx)
	return model, err
}

func InsertModel(ctx context.Context, model any) error {
	_, err := db.NewInsert().Model(model).Exec(ctx)
	return err
}

func UpdateWherePK(ctx context.Context, model any) error {
	_, err := db.NewUpdate().Model(model).WherePK().Exec(ctx)
	return err
}

func DeleteWherePK(ctx context.Context, model any) error {
	_, err := db.NewDelete().Model(model).WherePK().Exec(ctx)
	return err
}

func DeleteWherePKAndReturnID(ctx context.Context, model any) (returningID int64, err error) {
	_, err = db.NewDelete().Model(model).WherePK().Returning("id").Exec(ctx, &returningID)
	return
}

func GetWherePK[T any](ctx context.Context, id int64) (*T, error) {
	result := new(T)
	if err := db.NewSelect().Model(result).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return result, nil
}

func GetFromTable[T any](ctx context.Context, table string) (*[]T, error) {
	var result []T
	err := db.NewSelect().Table(table).Scan(ctx, &result)
	return &result, err
}

func GetFromModel[T any](ctx context.Context) (*[]T, error) {
	result := make([]T, 0)
	err := db.NewSelect().Model(&result).Scan(ctx)
	return &result, err
}
