package commands

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func HandleCommand(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "create":
		CreateRelationship(ctx, s, i)
	case "remove":
		RemoveRelationship(ctx, s, i)
	case "assign":
		AssignRole(ctx, s, i)
	case "unassign":
		UnassignRole(ctx, s, i)
	}
}
