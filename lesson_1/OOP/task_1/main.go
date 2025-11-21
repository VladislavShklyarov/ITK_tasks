package main

import "fmt"

func main() {
	kia := NewCar("Kia")
	fmt.Println(kia.Honk())

	volvoTruck := NewTruck("Volvo", 5)
	fmt.Println(volvoTruck.Honk())

	tesla := NewElectricCar("Tesla", 100)
	nissan := NewElectricCar("Nissan", 3)

	vehicles := []Vehicle{kia, volvoTruck, tesla, nissan}

	for _, vehicle := range vehicles {
		fmt.Println("Информация:", vehicle.GetInfo())

		// Запуск двигателя
		err := vehicle.StartEngine()
		if err != nil {
			fmt.Println("Ошибка запуска:", err)
		} else {
			fmt.Println("Двигатель успешно запущен")
		}

		fmt.Println("После запуска:", vehicle.GetInfo())

		// Повторный запуск (должен быть ошибка)
		err = vehicle.StartEngine()
		if err != nil {
			fmt.Println("Ошибка повторного запуска:", err)
		}

		// Остановка двигателя
		err = vehicle.StopEngine()
		if err != nil {
			fmt.Println("Ошибка остановки:", err)
		} else {
			fmt.Println("Двигатель успешно остановлен")
		}

		fmt.Println("После остановки:", vehicle.GetInfo())
		fmt.Println()
	}

	// Тестирование уникальных методов
	fmt.Println("\n=== Уникальные методы ===")
	fmt.Println("Kia honk:", kia.Honk())
	fmt.Println("Volvo honk:", volvoTruck.Honk())
	fmt.Println("Tesla battery:", tesla.GetBatteryLevel())
	fmt.Println("Volvo capacity:", volvoTruck.GetCargoCapacity())
}
