package smartwatch

import (
	"fmt"
	ce "task_3/customErrors"
	"unicode/utf8"
)

type Smartwatch struct {
	Brand        string
	Model        string
	SerialNumber string
	Year         int
	OS           *OSWatch
}

func NewSmartwatch(
	Brand string,
	Model string,
	SerialNumber string,
	Year int,
	OS *OSWatch,
) *Smartwatch {
	return &Smartwatch{
		Brand:        Brand,
		Model:        Model,
		SerialNumber: SerialNumber,
		Year:         Year,
		OS:           OS,
	}
}

func (sw *Smartwatch) UpdateOS(version string) error {
	if utf8.RuneCountInString(version) < 5 {
		return ce.ErrUnsupported
	}

	sw.OS.SetVersion(version)
	return nil
}

func (sw *Smartwatch) GetInfo() string {
	return fmt.Sprintf("Brand: %s, Model: %s, Year: %d, OS: %s %s",
		sw.Brand,
		sw.Model,
		sw.Year,
		sw.OS.GetName(),
		sw.OS.GetVersion())
}
