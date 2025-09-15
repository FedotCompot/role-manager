package database

import (
	"context"
	"role-manager-bot/internal/models"
)

func SetGuildSetting(ctx context.Context, guildID string, setting models.GuildSetting, value any) error {
	_, err := db.NewInsert().
		Model(&models.GuildSettings{GuildID: guildID, Settings: map[models.GuildSetting]any{setting: value}}).
		On("CONFLICT (guild_id) DO UPDATE").
		Set("settings = guild_settings.settings || EXCLUDED.settings").
		Exec(ctx)
	return err
}

func GetGuildSettings(ctx context.Context, guildID string) (map[models.GuildSetting]any, error) {
	guildSettings := models.GuildSettings{GuildID: guildID}
	err := db.NewSelect().
		Model(&guildSettings).
		WherePK().
		Scan(ctx)

	return guildSettings.Settings, err
}

func UpdateGuildDefaultSettings(ctx context.Context, guildID string) error {
	_, err := db.NewRaw(`
	INSERT INTO guild_settings (guild_id, settings)
	VALUES (?, ?)
	ON CONFLICT (guild_id) DO UPDATE
	SET settings = EXCLUDED.settings || guild_settings.settings`, guildID, models.GuildSettingsDefaults).Exec(ctx)
	return err
}

func UpdateDefaultSettings(ctx context.Context) error {
	_, err := db.NewRaw(`
	UPDATE guild_settings
	SET settings = ?::jsonb || settings`, models.GuildSettingsDefaults).Exec(ctx)
	return err
}
