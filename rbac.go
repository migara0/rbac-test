package main

import (
	"slices"
)

var roles = map[string][]string{
	"admin":    {"manage_users", "manage_order", "manage_products", "view_reports"},
	"seller":   {"create_product", "edit_products", "delete_product", "view_orders"},
	"customer": {"view_products", "place_order"},
}

type User struct {
	Username string
	Role     string
}

func (u User) GetPermissions() []string {
	return roles[u.Role]
}

func HasPermission(user User, action string) bool {
	return slices.Contains(user.GetPermissions(), action)
}
