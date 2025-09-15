package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewRaw(
			`CREATE TABLE guild_settings(
			   guild_id VARCHAR(32),
			   settings jsonb,
			   PRIMARY KEY (guild_id)
			)`).
			Exec(ctx)
		if err != nil {
			return err
		}
		return err
	}, func(ctx context.Context, db *bun.DB) error { return nil })
}
