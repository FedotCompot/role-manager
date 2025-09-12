package commands

import (
	"context"
	"fmt"
	"log/slog"
	"role-manager-bot/internal/database"

	"github.com/bwmarrin/discordgo"
)

func AssignRole(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	err = s.GuildMemberRoleAdd(i.GuildID, user.ID, role.ID)
	if err != nil {
		slog.Error("Failed to assign role.", "error", err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Failed to assign role.",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Assigned <@&%s> to <@%s>", role.ID, user.ID),
		},
	})
}
