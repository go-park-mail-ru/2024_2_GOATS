package redis

import (
	"fmt"
	"os"
)

func Init() (string, error) {
	// Получаем значения переменных окружения
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	//// Если переменные окружения не заданы, используем значения по умолчанию
	//if host == "" {
	//	host = "localhost"
	//}
	//if port == "" {
	//	port = "6379"
	//}

	addr := fmt.Sprintf("%s:%s", host, port)

	return addr, nil
}
