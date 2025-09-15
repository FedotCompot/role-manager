package features

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"reflect"
	"role-manager-bot/internal/database"
	"role-manager-bot/internal/models"

	"github.com/bwmarrin/discordgo"
)

const defaultChance float64 = 1

func randomReactions(ctx context.Context, s *discordgo.Session, i *discordgo.MessageCreate) {
	slog.Info("Random reaction")
	if i.Member == nil {
		return
	}
	slog.Info("Random reaction")
	emojiChance := defaultChance
	settings, err := database.GetGuildSettings(ctx, i.GuildID)
	if err == nil {
		slog.Info("DB chance", "chance", settings[models.SETTING_RANDOM_REACTION_CHANCE], "type", reflect.TypeOf(settings[models.SETTING_RANDOM_REACTION_CHANCE]))
		if dbChance, ok := settings[models.SETTING_RANDOM_REACTION_CHANCE].(float64); ok {
			emojiChance = dbChance
		}
	}
	slog.Info("Random reaction", "chance", emojiChance)
	if rand.Float64()*100 > emojiChance {
		return
	}
	emojis, err := s.GuildEmojis(i.GuildID)
	slog.Info("Random reaction chance proc", "emojis", emojis, "err", err)
	if err != nil || len(emojis) == 0 {
		return
	}
	err = s.MessageReactionAdd(i.ChannelID, i.Message.ID, emojis[rand.IntN(len(emojis))].APIName())
	slog.Error("Failed adding reaction", "error", err)
}
