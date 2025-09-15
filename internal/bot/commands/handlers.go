package commands

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func HandleCommand(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "This command must be used in a server",
			},
		})
		return
	}
	switch i.ApplicationCommandData().Name {
	case "create":
		CreateRelationship(ctx, s, i)
	case "remove":
		RemoveRelationship(ctx, s, i)
	case "assign":
		AssignRole(ctx, s, i)
	case "unassign":
		UnassignRole(ctx, s, i)
	case "settings":
		GuildSettings(ctx, s, i)
	}
}
