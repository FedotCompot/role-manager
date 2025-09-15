package database

import (
	"context"
	"fmt"
	"role-manager-bot/internal/bot/utils"
	"role-manager-bot/internal/models"

	"github.com/bwmarrin/discordgo"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func IsManager(ctx context.Context, member *discordgo.Member, role *discordgo.Role) (bool, error) {
	if utils.IsAdmin(member) {
		return true, nil
	}
	var rolePairs []*models.RoleRoleManager
	fmt.Println(member.Roles)
	err := db.NewSelect().
		Model(&rolePairs).
		Where("parent_role = ANY (?)", pgdialect.Array(member.Roles)).
		Scan(ctx)
	if err != nil {
		return false, err
	}

	for _, pair := range rolePairs {
		if role.ID == pair.ChildRole {
			return true, nil
		}
	}

	var userRolePairs []*models.UserRoleManager
	err = db.NewSelect().
		Model(&userRolePairs).
		Where("user_id = ?", member.User.ID).
		Scan(ctx)
	if err != nil {
		return false, err
	}

	for _, pair := range userRolePairs {
		if role.ID == pair.ChildRole {
			return true, nil
		}
	}

	return false, nil
}
