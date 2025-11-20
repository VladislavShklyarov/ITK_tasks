package smartphone

type OSMobile struct {
	name    string
	version string
}

func NewOSMobile(name string, version string) *OSMobile {
	return &OSMobile{
		name:    name,
		version: version,
	}
}

func (os *OSMobile) SetVersion(version string) {
	os.version = version
}
func (os *OSMobile) GetVersion() string {
	return os.version
}
func (os *OSMobile) GetName() string {
	return os.name
}
