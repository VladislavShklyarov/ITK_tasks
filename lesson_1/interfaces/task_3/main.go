package main

import (
	"fmt"
	"main/devices/laptop"
	"main/devices/smartphone"
	watch "main/devices/smartwatch"
)

func main() {

	IOS20 := smartphone.NewOSMobile("IOS", "11") // Версия ниже 12

	iphone := smartphone.NewSmartphone(
		"Apple",
		"13 Pro",
		"X800084YU7I31",
		2023,
		IOS20,
		3,
		256,
	)

	err := iphone.UpdateOS("20") // Обновляем до 20 версии
	if err == nil {
		fmt.Println(iphone.GetInfo()) // Ошибки нет, т.к. 11 < 12
	}

	err = iphone.UpdateOS("21") // Пробуем обновить до 21 версии
	if err != nil {
		fmt.Println(iphone.GetInfo()) // Ошибка будет, т.к. 20 уже больше 12
	}

	fmt.Println(IOS20.GetVersion()) // Объект обновляется.

	macOS := laptop.NewOSLaptop("macOS", "19", "x64")
	macbook := laptop.NewLaptop("Apple", "Air", "Y700003IU2", 2020, macOS)
	// Создали макбук

	err = macbook.UpdateOS("20")

	if err != nil {
		fmt.Println(macbook.GetInfo())
	}

	windows7 := laptop.NewOSLaptop("windows", "7", "x64")
	asus := laptop.NewLaptop("Asus", "Zenbook", "X800000q3", 2021, windows7)
	fmt.Println(asus.GetInfo())
	err = asus.UpdateOS("8")

	if err == nil {
		fmt.Println(asus.GetInfo()) // Asus обновился
	}

	watchOS := watch.NewOSWatch("watchOS", "10.2.5")
	appleWatch := watch.NewSmartwatch("Apple", "Watch 6", "Z90000x15", 2022, watchOS)
	fmt.Println(appleWatch.GetInfo())
	err = appleWatch.UpdateOS("10.3")
	if err != nil {
		fmt.Println(err) // Короче 5 символов
	}
	err = appleWatch.UpdateOS("10.3.6")
	fmt.Println(appleWatch.GetInfo())

}
