package banks

import (
	ce "main/customErrors"
	"main/utils"
)

func ConnectBank(bankName, APIKey string) (PaymentProcessor, error) {
	switch bankName {
	case "sberbank":

		if utils.CheckAPIKey(bankName, APIKey) == nil {
			return ConnectSberbank()
		} else {
			return nil, ce.ErrWrongAPIKey
		}
	case "tbank":
		if utils.CheckAPIKey(bankName, APIKey) == nil {
			return ConnectTbank()
		} else {
			return nil, ce.ErrWrongAPIKey
		}
	case "alfabank":
		if utils.CheckAPIKey(bankName, APIKey) == nil {
			return ConnectAlfa()
		} else {
			return nil, ce.ErrWrongAPIKey
		}
	default:
		return nil, ce.ErrWrongBank
	}
}
