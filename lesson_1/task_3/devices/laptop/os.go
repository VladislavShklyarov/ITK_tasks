package laptop

type OSLaptop struct {
	name    string
	version string
	arch    string
}

func NewOSLaptop(name string, version string, Arch string) *OSLaptop {
	return &OSLaptop{
		name:    name,
		version: version,
		arch:    Arch,
	}
}

func (os *OSLaptop) SetVersion(version string) {
	os.version = version
}
func (os *OSLaptop) GetVersion() string {
	return os.version
}
func (os *OSLaptop) GetName() string {
	return os.name
}
func (os *OSLaptop) GetArch() string {
	return os.arch
}
