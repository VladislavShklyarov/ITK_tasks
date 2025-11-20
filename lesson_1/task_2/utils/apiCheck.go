package utils

import (
	"os"
	ce "task_2/customErrors"
)

func CheckAPIKey(bankName, APIKey string) error {
	key := bankName + "APIKey"
	stored := os.Getenv(key)

	if stored != APIKey {
		return ce.ErrWrongAPIKey
	}
	return nil
}
