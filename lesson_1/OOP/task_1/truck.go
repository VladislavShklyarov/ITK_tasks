package main

import "fmt"

type Truck struct {
	Car
	CargoCapacity int
}

func (t *Truck) Honk() string {
	return "Honk Honk!"
}

func NewTruck(Brand string, capacity int) *Truck {
	return &Truck{
		Car:           Car{Brand: Brand},
		CargoCapacity: capacity,
	}
}

func (t *Truck) StartEngine() error {
	if t.GetEngineStatus() == "запущен" {
		return fmt.Errorf("двигатель уже запущен")
	} else {
		t.engineOn = true
	}
	return nil
}

func (t *Truck) StopEngine() error {
	if t.GetEngineStatus() == "заглушен" {
		return fmt.Errorf("двигатель уже заглушен")
	} else {
		t.engineOn = false
	}
	return nil
}

func (t *Truck) GetEngineStatus() string {
	if t.engineOn {
		return "запущен"
	}
	return "заглушен"
}

func (t *Truck) GetCargoCapacity() int {
	return t.CargoCapacity
}

func (t *Truck) GetInfo() string {
	engineStatus := t.GetEngineStatus()

	return fmt.Sprintf("Марка: %s, Грузоподъемность: %d тонн, Двигатель: %s",
		t.Brand, t.CargoCapacity, engineStatus)
}
