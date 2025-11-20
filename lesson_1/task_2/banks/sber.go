package banks

import (
	"fmt"
	"math/rand"
	ce "task_2/customErrors"
)

type Sberbank struct{}

func ConnectSberbank() (*Sberbank, error) {
	fmt.Println("connecting Sber...")

	if rand.Intn(20) == 0 { // Сбер отвалится в 5% случаев
		return nil, ce.ErrProviderUnavailable
	}
	return &Sberbank{}, nil
}

func (s *Sberbank) ProcessPayment(amount float64) error {

	if amount <= 0 {
		return ce.ErrInvalidAmount
	}
	fmt.Println("Успешно. Спасибо, что используете Тбанк! ")
	return nil
}
