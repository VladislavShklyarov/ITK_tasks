package customErrors

import "errors"

var (
	ErrUnsupported          = errors.New("обновление недоступно")
	ErrWrongMobileOSVersion = errors.New("версия мобильной ос может быть только в формате float64")
	ErrWrongLaptopOSType    = errors.New("обновления поддерживаются только ноутбуками на windows")
)
