package configManager

type configManager interface {
	LoadConfig()
	Get(key string) string
	PrintConfig()
}

// интерфейс для того, чтобы можно было реализовать также загрузку из БД, файла или еще откуда-нибудь.

type Config struct {
	Vault map[string]string
}
