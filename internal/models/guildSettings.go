package models

import "github.com/uptrace/bun"

type GuildSettings struct {
	bun.BaseModel `bun:"table:guild_settings"`

	GuildID  string               `bun:"guild_id,pk" json:"guild_id"`
	Settings map[GuildSetting]any `bun:"settings" json:"settings"`
}

type GuildSetting string

const (
	SETTING_RANDOM_REACTION_CHANCE GuildSetting = "random_reaction_chance"
)

var guildSettingsFromString map[string]GuildSetting = map[string]GuildSetting{
	"random_reaction_chance": SETTING_RANDOM_REACTION_CHANCE,
}
var GuildSettingsDefaults map[GuildSetting]any = map[GuildSetting]any{
	SETTING_RANDOM_REACTION_CHANCE: float64(1),
}

func GuildSettingFromSting(setting string) (GuildSetting, bool) {
	sett, ok := guildSettingsFromString[setting]
	return sett, ok
}
