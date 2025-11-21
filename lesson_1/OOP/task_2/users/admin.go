package users

import (
	"errors"
	"os"
)

type Admin struct {
	*Moderator
}

func NewAdmin(username string) *Admin {
	a := &Admin{
		NewModerator(username),
	}
	a.permissions.Add("delete")
	a.permissions.Add("manage_roles")
	return a
}

func (a *Admin) GetUsername() string {
	return a.username
}

func (a *Admin) GetRole() string {
	return "powerful admin!"
}

func (a *Admin) HasPermission(permission string) bool {
	return a.permissions.Has(permission)
}

func (a *Admin) ShowPermissions(password string) ([]string, error) { // чтобы посмотреть разрешения админа нужен пароль
	correctPassword := os.Getenv("adminPassword")
	if password == correctPassword {
		return a.permissions.GetPermissions(), nil
	} else {
		return nil, errors.New("wrong password for admin")
	}
}
