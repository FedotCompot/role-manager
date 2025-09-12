package commands

import (
	"context"
	"fmt"
	"log/slog"
	"role-manager-bot/internal/database"
	"slices"

	"github.com/bwmarrin/discordgo"
)

func AssignRole(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	user, err := s.GuildMember(i.GuildID, i.ApplicationCommandData().GetOption("user").UserValue(nil).ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "User must be a member of this server",
			},
		})
		return
	}
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
	if slices.Contains(user.Roles, role.ID) {

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "User already has this role",
			},
		})
		return
	}

	err = s.GuildMemberRoleAdd(i.GuildID, user.User.ID, role.ID)
	if err != nil {
		slog.Error("Failed to assign role.", "error", err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Failed to assign role, bot must have a higher role to manage it",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Assigned <@&%s> to <@%s>", role.ID, user.User.ID),
		},
	})
}
