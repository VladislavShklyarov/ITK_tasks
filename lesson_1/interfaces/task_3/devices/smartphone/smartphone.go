package smartphone

import (
	"fmt"
	ce "main/customErrors"
	"strconv"
)

type Smartphone struct {
	Brand         string
	Model         string
	SerialNumber  string
	Year          int
	OS            *OSMobile
	CamerasAmount int
	Memory        int
}

func NewSmartphone(
	Brand string,
	Model string,
	SerialNumber string,
	Year int,
	OS *OSMobile,
	CameraAmount int,
	Memory int,
) *Smartphone {
	return &Smartphone{
		Brand:         Brand,
		Model:         Model,
		SerialNumber:  SerialNumber,
		Year:          Year,
		OS:            OS,
		CamerasAmount: CameraAmount,
		Memory:        Memory,
	}
}

func (s *Smartphone) UpdateOS(version string) error {

	_, err := strconv.Atoi(version)
	if err != nil {
		return ce.ErrWrongMobileOSVersion
	}

	osVersion, _ := strconv.Atoi(s.OS.GetVersion()) // лучше приводить к типу float
	if osVersion >= 12.0 {
		return ce.ErrUnsupported
	}
	s.OS.SetVersion(version)
	return nil

}

func (s *Smartphone) GetInfo() string {
	return fmt.Sprintf("Brand: %s, Model: %s, Year: %d, OS: %s%s, Cameras: %d, Memory: %dGB",
		s.Brand,
		s.Model,
		s.Year,
		s.OS.GetName(),
		s.OS.GetVersion(),
		s.CamerasAmount,
		s.Memory)
}
