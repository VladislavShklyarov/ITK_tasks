package utils

import (
	ce "main/customErrors"
	"os"
)

func CheckAPIKey(bankName, APIKey string) error {
	key := bankName + "APIKey"
	stored := os.Getenv(key)

	if stored != APIKey {
		return ce.ErrWrongAPIKey
	}
	return nil
}
