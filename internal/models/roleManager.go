package models

import "github.com/uptrace/bun"

type RoleRoleManager struct {
	bun.BaseModel `bun:"table:role_role_manager"`
	ParentRole    string `bun:"parent_role"`
	ChildRole     string `bun:"child_role"`
}

type UserRoleManager struct {
	bun.BaseModel `bun:"table:user_role_manager"`
	ParentUser    string `bun:"user_id"`
	ChildRole     string `bun:"child_role"`
}
