package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/banks"
	"math/rand"
	"time"
)

func init() {
	err := godotenv.Load("./cfg/.env")
	if err != nil {
		log.Fatal("Не удалось загрузить .env файл")
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	sber, err := banks.ConnectBank("sberbank", "123")
	if err != nil {
		fmt.Println(err)
	} else {
		err = sber.ProcessPayment(10)
		if err != nil {
			fmt.Println(err)
		}
	}

	tbank, err := banks.ConnectBank("tbank", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		err = tbank.ProcessPayment(0)
		if err != nil {
			fmt.Println(err)
		}
	}

	alfabank, err := banks.ConnectBank("alfabank", "789")

	if err != nil {
		fmt.Println(err)
	} else {
		err = alfabank.ProcessPayment(12)
		if err != nil {
			fmt.Println(err)
		}
	}
}
