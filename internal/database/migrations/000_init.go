package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewRaw(
			`CREATE TABLE role_role_manager(
			   parent_role VARCHAR(32),
			   child_role VARCHAR(32),
			   PRIMARY KEY (parent_role, child_role)
			)`).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewRaw(
			`CREATE TABLE user_role_manager(
				   user_id VARCHAR(32),
				   child_role VARCHAR(32),
				   PRIMARY KEY (user_id, child_role)
				)`).
			Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error { return nil })
}
