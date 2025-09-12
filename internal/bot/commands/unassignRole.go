package commands

import (
	"context"
	"fmt"
	"role-manager-bot/internal/database"

	"github.com/bwmarrin/discordgo"
)

func UnassignRole(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().GetOption("user").UserValue(nil)
	role := i.ApplicationCommandData().GetOption("role").RoleValue(nil, "")

	if user == nil || role == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Both user and role are required",
			},
		})
		return
	}

	isParent, err := database.IsManager(ctx, i.Member, role)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Server error",
			},
		})
		return
	}

	if !isParent {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "You do not have permission to manage this role.",
			},
		})
		return
	}

	err = s.GuildMemberRoleRemove(i.GuildID, user.ID, role.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Failed to unassign role.",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Unassigned <@&%s> from <@%s>", role.ID, user.ID),
		},
	})
}
