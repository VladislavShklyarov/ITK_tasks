package main

import "fmt"

type ElectricCar struct {
	Car
	BatteryLevel int
}

func NewElectricCar(Brand string, batteryLevel int) *ElectricCar {
	return &ElectricCar{
		Car:          Car{Brand: Brand},
		BatteryLevel: batteryLevel,
	}
}

func (e *ElectricCar) StartEngine() error {
	if e.GetBatteryLevel() <= 5 {
		return fmt.Errorf("низкий заряд батареи")
	} else if e.GetEngineStatus() == "запущен" {
		return fmt.Errorf("двигатель уже запущен")
	} else {
		e.engineOn = true
	}
	return nil
}

func (e *ElectricCar) StopEngine() error {
	if e.engineOn == false {
		return fmt.Errorf("двигатель уже заглушен")
	} else {
		e.engineOn = false
	}
	return nil
}

func (e *ElectricCar) GetEngineStatus() string {
	if e.engineOn {
		return "запущен"
	}
	return "заглушен"
}

func (e *ElectricCar) GetBatteryLevel() int {
	return e.BatteryLevel
}

func (e *ElectricCar) GetInfo() string {
	engineStatus := e.GetEngineStatus()
	batteryLevel := e.GetBatteryLevel()

	return fmt.Sprintf("Марка: %s, Заряд батареи: %d%%, Двигатель: %s",
		e.Brand, batteryLevel, engineStatus)
}
