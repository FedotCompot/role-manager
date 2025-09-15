package commands

import (
	"context"
	"fmt"
	"log/slog"
	"role-manager-bot/internal/bot/utils"
	"role-manager-bot/internal/database"
	"role-manager-bot/internal/models"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GuildSettings(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !utils.IsAdmin(i.Member) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "This command can be used only by an Administrator",
			},
		})
		return
	}

	setSubcommand := i.ApplicationCommandData().GetOption("set")
	if setSubcommand != nil {
		setting := setSubcommand.Options[0]
		value := setting.GetOption("value").Value
		guildSetting, ok := models.GuildSettingFromSting(setting.Name)
		if !ok {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "This setting does not exist",
				},
			})
			return
		}
		err := database.SetGuildSetting(ctx, i.GuildID, guildSetting, value)

		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Failed to update setting",
				},
			})
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Setting updated!",
			},
		})
	}

	getSubcommand := i.ApplicationCommandData().GetOption("get")
	if getSubcommand != nil {
		settings, err := database.GetGuildSettings(ctx, i.GuildID)

		if err != nil {
			slog.Error("Failed getting server settings", "error", err)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Failed getting server settings",
				},
			})
			return
		}
		slog.Info("Settings", "value", settings)
		toSend := &strings.Builder{}
		toSend.WriteString("Settings:\n\n")
		for setting, value := range settings {
			fmt.Fprintf(toSend, "%s: %#v\n", setting, value)
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: toSend.String(),
			},
		})
	}
}
