package database

import (
	"context"
	"fmt"
	"role-manager-bot/internal/models"

	"github.com/bwmarrin/discordgo"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func IsManager(ctx context.Context, user *discordgo.Member, role *discordgo.Role) (bool, error) {
	if user.Permissions&discordgo.PermissionAdministrator != 0 {
		return true, nil
	}
	var rolePairs []*models.RoleRoleManager
	fmt.Println(user.Roles)
	err := db.NewSelect().
		Model(&rolePairs).
		Where("parent_role = ANY (?)", pgdialect.Array(user.Roles)).
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
		Where("user_id = ?", user.User.ID).
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
