package laptop

import (
	"fmt"
	ce "main/customErrors"
)

type Laptop struct {
	Brand        string
	Model        string
	SerialNumber string
	Year         int
	OS           *OSLaptop
}

func NewLaptop(
	Brand string,
	Model string,
	SerialNumber string,
	Year int,
	OS *OSLaptop,
) *Laptop {
	return &Laptop{
		Brand:        Brand,
		Model:        Model,
		SerialNumber: SerialNumber,
		Year:         Year,
		OS:           OS,
	}
}
func (l *Laptop) UpdateOS(version string) error {
	if l.OS.GetName() != "windows" {
		return ce.ErrWrongLaptopOSType
	}

	l.OS.SetVersion(version)
	return nil

}
func (l *Laptop) GetInfo() string {
	return fmt.Sprintf("Brand: %s, Model: %s, Year: %d, OS: %s %s (%s), Serial: %s",
		l.Brand,
		l.Model,
		l.Year,
		l.OS.GetName(),
		l.OS.GetVersion(),
		l.OS.GetArch(),
		l.SerialNumber)
}
