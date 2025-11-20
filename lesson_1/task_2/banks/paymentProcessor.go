package banks

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}
