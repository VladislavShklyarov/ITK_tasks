package banks

import (
	"fmt"
	"math/rand"
	ce "task_2/customErrors"
)

type Tbank struct{}

func (t *Tbank) ProcessPayment(amount float64) error {

	if amount <= 0 {
		return ce.ErrInvalidAmount
	}
	fmt.Println("Успешно. Спасибо, что используете Тбанк! ")
	return nil
}

func ConnectTbank() (*Tbank, error) {
	fmt.Println("connecting Tbank...")

	if rand.Intn(10) == 0 { // Тбанк отвалится в 10% случаев
		return nil, ce.ErrProviderUnavailable
	}

	return &Tbank{}, nil
}
