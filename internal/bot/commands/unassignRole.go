package commands

import (
	"context"
	"fmt"
	"role-manager-bot/internal/database"

	"github.com/bwmarrin/discordgo"
)

func UnassignRole(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	member := i.Member
	user := options[0].UserValue(nil)
	role := options[1].RoleValue(nil, "")

	if member == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This command must be used in a server",
			},
		})
		return
	}
	if user == nil || role == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Both user and role are required",
			},
		})
		return
	}

	isParent, err := database.IsParent(ctx, member, role)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Server error",
			},
		})
		return
	}

	if !isParent {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
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
