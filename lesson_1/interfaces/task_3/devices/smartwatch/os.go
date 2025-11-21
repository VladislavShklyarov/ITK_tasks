package smartwatch

type OSWatch struct {
	name    string
	version string
}

func NewOSWatch(name string, version string) *OSWatch {
	return &OSWatch{
		name:    name,
		version: version,
	}
}

func (os *OSWatch) SetVersion(version string) {
	os.version = version
}

func (os *OSWatch) GetVersion() string {
	return os.version
}

func (os *OSWatch) GetName() string {
	return os.name
}
