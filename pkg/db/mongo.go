package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect устанавливает соединение с базой данных MongoDB
func Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Замените на ваш URI подключения

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("ошибка подключения к MongoDB: %v", err)
	}

	// Проверяем подключение
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("ошибка ping к MongoDB: %v", err)
	}

	log.Println("Успешно подключено к MongoDB!")
	return nil
}

// Disconnect закрывает соединение с базой данных MongoDB
func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Ошибка отключения от MongoDB: %v", err)
	}
	log.Println("Отключено от MongoDB.")
}

// GetCollection получает коллекцию из базы данных
func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("football_simulator").Collection(collectionName) // Измените имя базы данных, если нужно
}
