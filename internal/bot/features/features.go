package features

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func Init(ctx context.Context, s *discordgo.Session) {
	s.AddHandler(RandomReactions)
}
