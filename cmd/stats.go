package cmd

import (
	"Footballsim/pkg/db" // Импортируем пакет для работы с MongoDB
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// MatchStats представляет статистику одного матча
type MatchStats struct {
	Team1    string `json:"team1"`
	Team2    string `json:"team2"`
	Score1   int    `json:"score1"`
	Score2   int    `json:"score2"`
	Duration int    `json:"duration"`
}

// FetchMatchStats извлекает статистику матчей из базы данных
func FetchMatchStats() ([]MatchStats, error) {
	collection := db.GetCollection("matches")
	findOptions := options.Find()
	var matchStats []MatchStats

	// Поиск всех матчей
	cursor, err := collection.Find(context.TODO(), nil, findOptions)
	if err != nil {
		return nil, fmt.Errorf("ошибка поиска матчей: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Проход по результатам поиска и добавление их в срез matchStats
	for cursor.Next(context.TODO()) {
		var match MatchStats
		if err := cursor.Decode(&match); err != nil {
			return nil, fmt.Errorf("ошибка декодирования матча: %v", err)
		}
		matchStats = append(matchStats, match)
	}

	// Проверка на ошибки после итерации
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при переборе курсора: %v", err)
	}

	return matchStats, nil
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Display statistics of past matches",
	Long:  `This command displays the statistics of previously simulated football matches.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Displaying match statistics...")
		matchStats, err := FetchMatchStats()
		if err != nil {
			log.Fatalf("Ошибка при получении статистики матчей: %v", err)
		}

		// Выводим статистику матчей
		for _, match := range matchStats {
			fmt.Printf("Матч: %s %d:%d %s (Продолжительность: %d минут)\n",
				match.Team1, match.Score1, match.Score2, match.Team2, match.Duration)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
