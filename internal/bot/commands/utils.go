package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func Roll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	min := uint64(0)
	max := uint64(100)

	if minOption := i.ApplicationCommandData().GetOption("min"); minOption != nil {
		min = minOption.UintValue()
	}
	if maxOption := i.ApplicationCommandData().GetOption("max"); maxOption != nil {
		max = maxOption.UintValue()
	}
	if min > max {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Min must be < max",
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Roll (%d-%d): %d", min, max, rand.Uint64N(max+1-min)+min),
		},
	})
}
