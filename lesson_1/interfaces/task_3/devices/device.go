package devices

type Device interface {
	UpdateOS(version string) error
	GetInfo() string
}
