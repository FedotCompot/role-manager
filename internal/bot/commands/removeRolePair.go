package commands

import (
	"context"
	"fmt"
	"role-manager-bot/internal/database"
	"role-manager-bot/internal/models"

	"github.com/bwmarrin/discordgo"
)

func RemoveRelationship(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	subCommandRole := i.ApplicationCommandData().GetOption("role")
	subCommandUser := i.ApplicationCommandData().GetOption("user")

	if subCommandRole != nil {
		parentRole := subCommandRole.GetOption("parent").RoleValue(nil, "")
		childRole := subCommandRole.GetOption("child").RoleValue(nil, "")
		err := database.DeleteWherePK(ctx, &models.RoleRoleManager{ParentRole: parentRole.ID, ChildRole: childRole.ID})
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Failed to remove role relationship.",
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Role relationship removed: <@&%s> is no longer a manager of <@&%s>", parentRole.ID, childRole.ID),
			},
		})
	} else if subCommandUser != nil {
		parentUser := subCommandUser.GetOption("user").UserValue(nil)
		childRole := subCommandUser.GetOption("child").RoleValue(nil, "")
		err := database.DeleteWherePK(ctx, &models.UserRoleManager{ParentUser: parentUser.ID, ChildRole: childRole.ID})
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Failed to remove user relationship.",
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("User-role relationship removed: <@%s> is no longer manager of <@&%s>", parentUser.ID, childRole.ID),
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Parent user or role has to be provided",
			},
		})
	}
}
