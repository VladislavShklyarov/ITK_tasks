package users

import (
	ps "main/permissions"
)

type BasicUser struct {
	username    string
	permissions ps.Permissions
}

func NewBasicUser(username string) *BasicUser {
	u := &BasicUser{
		username:    username,
		permissions: ps.NewPermissions(),
	}
	u.permissions.Add("read")
	return u
}

func (bu *BasicUser) GetUsername() string {
	return bu.username
}

func (bu *BasicUser) GetRole() string {
	return "basic user"
}

func (bu *BasicUser) HasPermission(permission string) bool {
	return bu.permissions.Has(permission)
}

func (bu *BasicUser) ShowPermissions() []string {
	return bu.permissions.GetPermissions()
}
