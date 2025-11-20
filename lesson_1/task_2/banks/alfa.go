package banks

import (
	"fmt"
	"math/rand"
	ce "task_2/customErrors"
)

func (a *Alfabank) ProcessPayment(amount float64) error {

	if amount <= 0 {
		return ce.ErrInvalidAmount
	}
	fmt.Println("Успешно. Спасибо, что используете Альфабанк! ")
	return nil
}

type Alfabank struct{}

func ConnectAlfa() (*Alfabank, error) {
	fmt.Println("connecting Alfabank...")

	if rand.Intn(2) == 0 { // Алфьабанк отвалится в 50% случаев
		return nil, ce.ErrProviderUnavailable
	}
	return &Alfabank{}, nil
}
