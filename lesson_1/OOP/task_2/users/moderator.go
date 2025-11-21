package users

type Moderator struct {
	*BasicUser
}

func NewModerator(username string) *Moderator {
	m := &Moderator{
		BasicUser: NewBasicUser(username),
	}
	m.permissions.Add("edit")
	m.permissions.Add("ban_users")
	return m
}

func (m *Moderator) GetUsername() string {
	return m.username
}

func (m *Moderator) GetRole() string {
	return "moderator"
}

func (m *Moderator) HasPermission(permission string) bool {
	return m.permissions.Has(permission)
}

func (m *Moderator) ShowPermissions() []string {
	return m.permissions.GetPermissions()
}
