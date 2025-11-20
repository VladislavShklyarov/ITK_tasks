package customErrors

import "errors"

var (
	ErrInvalidAmount       = errors.New("некорректная сумма платежа")
	ErrProviderUnavailable = errors.New("провайдер недоступен")
	ErrWrongAPIKey         = errors.New("неверный ключ API")
	ErrWrongBank           = errors.New("этот банк недоступен")
)
