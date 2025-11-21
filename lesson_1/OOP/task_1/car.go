package main

import "fmt"

type Car struct {
	Brand    string
	engineOn bool
}

func (c *Car) Honk() string {
	return "beep beep!"
}

func NewCar(Brand string) *Car {
	return &Car{Brand: Brand}
}

func (c *Car) StartEngine() error {
	if c.GetEngineStatus() == "запущен" {
		return fmt.Errorf("двигатель уже запущен")
	} else {
		c.engineOn = true
	}
	return nil
}

func (c *Car) StopEngine() error {
	if c.GetEngineStatus() == "заглушен" {
		return fmt.Errorf("двигатель уже заглушен")
	} else {
		c.engineOn = false
	}
	return nil
}

func (c *Car) GetEngineStatus() string {
	if c.engineOn {
		return "запущен"
	}
	return "заглушен"
}

func (c *Car) GetInfo() string {
	engineStatus := c.GetEngineStatus()

	return fmt.Sprintf("Марка: %s, Двигатель: %s",
		c.Brand, engineStatus)
}
