package main

import (
	"Footballsim/cmd"
	"Footballsim/pkg/db" // Импортируем пакет для работы с MongoDB
	"log"
)

func main() {
	// Подключение к MongoDB
	err := db.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	// Запуск приложения
	cmd.Execute()

	// Отключение от MongoDB после завершения выполнения команд
	db.Disconnect()

}
