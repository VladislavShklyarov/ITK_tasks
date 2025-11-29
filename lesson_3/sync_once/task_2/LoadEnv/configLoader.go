package LoadEnv

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
	"task_2/configManager"
)

type EnvConfLoader struct {
	once   sync.Once
	loaded *configManager.Config
}

func NewEnvConfigLoader() *EnvConfLoader {
	return &EnvConfLoader{
		once:   sync.Once{},
		loaded: &configManager.Config{Vault: make(map[string]string)},
	}
}

func (cl *EnvConfLoader) Get(key string) string {
	if value, ok := cl.loaded.Vault[key]; !ok {
		fmt.Println("такого ключа нет")
		return ""
	} else {
		return value
	}
}

func (cl *EnvConfLoader) PrintConfig() {
	if cl.loaded == nil {
		fmt.Println("Конфиг ещё не загружен, сначала загрузим .env")
		cl.LoadConfig()
	}
	for key, value := range cl.loaded.Vault {
		fmt.Printf("%s : %s\n", key, value)
	}
}

func (cl *EnvConfLoader) LoadConfig() {
	cl.once.Do(func() {
		envConfig, err := godotenv.Read("./cfg/.env")
		if err != nil {
			fmt.Println("Ошибка загрузки конфигурации из .env")
			return
		}
		cl.loaded.Vault = envConfig

		for key, val := range envConfig {
			os.Setenv(key, val)
		}
		fmt.Println("Конфиг загружен из .env")
	})
}
