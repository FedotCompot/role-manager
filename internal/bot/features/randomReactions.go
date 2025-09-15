package features

import (
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

const CHANCE = 100 // 1/CHANCE probability

func RandomReactions(s *discordgo.Session, i *discordgo.MessageCreate) {
	if i.Member == nil {
		return
	}
	if rand.UintN(CHANCE) > 0 {
		return
	}
	emojis, err := s.GuildEmojis(i.GuildID)
	if err != nil {
		return
	}
	s.MessageReactionAdd(i.ChannelID, i.Message.ID, emojis[rand.IntN(len(emojis))].ID)
}
