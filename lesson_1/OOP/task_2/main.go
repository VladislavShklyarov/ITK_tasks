package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/users"
	"math/rand"
	"time"
)

func init() {
	err := godotenv.Load("cfg/.env")
	if err != nil {
		log.Fatal("Не удалось загрузить .env файл")
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
func main() {
	basicUser := users.NewBasicUser("some_default_user")
	fmt.Println("Базовый пользователь:")
	fmt.Println("Имя:", basicUser.GetUsername())
	fmt.Println("Разрешение 'read':", basicUser.HasPermission("read"))
	fmt.Println("Разрешение 'edit':", basicUser.HasPermission("edit")) // должно быть false
	fmt.Println("Все разрешения:", basicUser.ShowPermissions())
	fmt.Println("Роль:", basicUser.GetRole())
	fmt.Println("===   ===")

	moderator := users.NewModerator("moderator")
	fmt.Println("Модератор:")
	fmt.Println("Имя:", moderator.GetUsername())
	fmt.Println("Разрешение 'read':", moderator.HasPermission("read"))     // унаследовано
	fmt.Println("Разрешение 'edit':", moderator.HasPermission("edit"))     // унаследовано
	fmt.Println("Разрешение 'delete':", moderator.HasPermission("delete")) // дополнительное право
	fmt.Println("Все разрешения:", moderator.ShowPermissions())
	fmt.Println("Роль:", moderator.GetRole())
	fmt.Println("===   ===")

	admin := users.NewAdmin("admin")
	fmt.Println("Администратор:")
	fmt.Println("Имя:", admin.GetUsername())
	fmt.Println("Разрешение 'read':", admin.HasPermission("read"))
	fmt.Println("Разрешение 'write':", admin.HasPermission("write"))
	fmt.Println("Разрешение 'delete':", admin.HasPermission("delete"))
	err, perms := admin.ShowPermissions("456")
	fmt.Println("Попытка показать права с неверным паролем:", err, perms)
	err, perms = admin.ShowPermissions("123")
	fmt.Println("Попытка показать права с правильным паролем:", err, perms)
}
