package bot

import (
	"context"
	"fmt"
	"log/slog"
	"role-manager-bot/internal/bot/commands"
	"role-manager-bot/internal/bot/features"
	"role-manager-bot/internal/config"

	"github.com/bwmarrin/discordgo"
)

var botContext context.Context

func Init(ctx context.Context) func() {
	var err error
	session, err := discordgo.New("Bot " + config.Data.DiscordBotToken)
	if err != nil {
		panic(err)
	}
	botContext = ctx
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMembers)
	session.AddHandler(ready)
	session.AddHandler(interactionCreate)
	features.Init(ctx, session)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	return func() {
		session.Close()
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "create",
			Description: "Create a parent-child role relationship",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "role",
					Description: "Create a parent-child role relationship",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "parent",
							Description: "Parent role",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "child",
							Description: "The child role",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "user",
					Description: "Create a user manager for child role",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionUser,
							Name:        "user",
							Description: "Role manager",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "child",
							Description: "The child role",
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:        "remove",
			Description: "Create a parent-child role relationship",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "role",
					Description: "Create a parent-child role relationship",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "parent",
							Description: "Parent role",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "child",
							Description: "The child role",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "user",
					Description: "Create a user manager for child role",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionUser,
							Name:        "user",
							Description: "Role manager",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "child",
							Description: "The child role",
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:        "assign",
			Description: "Assign a child role to a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to assign the role to",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The child role to assign",
					Required:    true,
				},
			},
		},
		{
			Name:        "unassign",
			Description: "Unassign a child role from a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to unassign the role from",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The child role to unassign",
					Required:    true,
				},
			},
		},
	}
	fmt.Println("Updating commands...")
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		slog.Error("Cannot create commands", "error", err)
	}
	fmt.Println("Bot is ready!")
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		commands.HandleCommand(botContext, s, i)
	}
}
