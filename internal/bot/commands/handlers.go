package commands

import (
	"context"
	"role-manager-bot/internal/bot/utils"

	"github.com/bwmarrin/discordgo"
)

func HandleCommand(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {

	switch i.ApplicationCommandData().Name {
	case "roll":
		Roll(s, i)
	}

	// Server-only commands
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

	// Public commands
	switch i.ApplicationCommandData().Name {
	case "assign":
		AssignRole(ctx, s, i)
	case "unassign":
		UnassignRole(ctx, s, i)
	}

	if !utils.IsAdmin(i.Member) {
		return
	}
	// Admin only commands
	switch i.ApplicationCommandData().Name {
	case "settings":
		GuildSettings(ctx, s, i)
	case "create":
		CreateRelationship(ctx, s, i)
	case "remove":
		RemoveRelationship(ctx, s, i)
	}
}
