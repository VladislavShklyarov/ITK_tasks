package permissions

type Permissions struct {
	perms map[string]bool
}

func NewPermissions() Permissions {
	return Permissions{
		perms: make(map[string]bool),
	}
}

func (p *Permissions) Add(permission string) {
	p.perms[permission] = true
}

func (p *Permissions) Has(permission string) bool {
	return p.perms[permission]
}

func (p *Permissions) GetPermissions() []string {
	perms := make([]string, 0, len(p.perms))

	for perm := range p.perms {
		perms = append(perms, perm)
	}
	return perms
}
