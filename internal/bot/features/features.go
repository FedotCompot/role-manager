package features

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func Init(ctx context.Context, s *discordgo.Session) {
	messageFeatures(ctx, s)
}

func messageFeatures(ctx context.Context, s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.MessageCreate) {
		randomReactions(ctx, s, i)
	})
}
