package utils

import "github.com/bwmarrin/discordgo"

func IsAdmin(member *discordgo.Member) bool {
	return member.Permissions&discordgo.PermissionAdministrator != 0
}
